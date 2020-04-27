# Copyright (c) 2020 PaddlePaddle Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from __future__ import print_function
from concurrent import futures
import data_server_pb2
import data_server_pb2_grpc
import master_pb2
import master_pb2_grpc
import grpc
import sys
import logging
from threading import Thread
import Queue
import DistributeReader
from exception import *
from dataset import EdlDataSet


class DataServerServicer(object):
    def __init__(self, master, data_set, capcity=1000):
        self._master = master
        self._sub_data_set = Queue()
        # {file_key:{rec_no: data}}
        self._data = {}
        # to control the cache size.
        self._data_queue = Queue(capcity)
        self._lock = threading.Lock

        assert isinstance(data_set, EdlDataSet)

        self._t_get_sub_dataset = Thread(target=self._get_sub_dataset)
        self._t_read_data = Thread(target=self._read_data)

    def _get_sub_dataset(self):
        channel = grpc.insecure_channel(self.master)
        stub = data_server_pb2_grpc.MasterStub(channel)
        request = data_server_pb2.SubDataSetRequest()
        response = stub.GetSubDataSet(request)
        for file_data_set in response.files:
            self._sub_data_set.put(file_data_set)

    def _get_file_key(self, idx, file_path):
        key = "idx:{}_path:{}".format(idx, file_path)
        return key

    def _read_data(self):
        while True:
            file_data_set = self._sub_data_set.get()

            rec_map = {}
            for rec in file_data_set.record:
                rec_map[rec.record_no] = rec.status
            for rec_no, data in self._data_set.read(file_data_set.file_path):
                if rec_map[rec_no] == RecordStatus.PROCSSED:
                    continue

                self._data_queue.put(1, block=True)
                key = self._get_file_key(file_data_set.idx_in_list,
                                         file_data_set.file_path)
                with self._lock:
                    if key not in self._data:
                        self._data[key] = {}
                    self._data[key][rec_no] = data

    def GetData(self, request, context):
        response = data_server_pb2.DataResponse()

        files = data_sever_pb2.Files()
        files_error = data_server_pb2.FilesError()

        for meta in request.metas:
            one_file = data_server_pb2.File()
            one_file.idx_in_list = meta.idx_in_list
            one_file.file_path = meta.file_path

            file_error = data_server_pb2.FileError()
            file_error.idx_in_list = meta.idx_in_list
            file_error.file_path = meta.file_path
            file_error.status = data_server_pb2.DataStatus.NOT_FOUND

            key = self._get_file_key(meta.idx_in_list, meta.file_path)
            while self._lock:
                if key not in self._data:
                    response.errors.append(file_error)
                    continue

                record = data_server_pb2.Record()
                record_error = data_server_pb2.RecordError()
                record_error.status = data_server_pb2.DataStatus.NOT_FOUND

                for rec_no in meta.record_no:
                    if rec_no not in self._data[key]:
                        record_error.rec_no = rec_no
                        file_error.errors.append(record_error)
                        response.errors.append(file_error)
                        continue

                    data = self._data[key][rec_no]

                    record.record_no = rec_no
                    record.data = data

                    one_file.records.append(record)
                if len(file_error.errors) > 0:
                    files_error.errors.append(file_error)
                files.files.append(one_file)

            if len(files_error) > 0:
                response.errors = files_error
                return response

        response.files = files
        return response

    def ClearDataCache(self, request, context):
        response = common.RPCRet()
        for meta in request.metas:
            file_key = self._get_file_key(meta.idx_in_list, meta.file_path)

            with self._lock():
                if file_key not in self._data[key]:
                    continue

                recs = self._data[key]
                for rec_no in recs:
                    if rec_no not in recs:
                        continue

                    recs.pop(rec_no)
                    self._data_queue.pop()

        return response


class DataServer(object):
    def __init__(self):
        pass

    def start(self, max_workers=1000, concurrency=100, endpoint=""):
        if endpoint == "":
            logging.error("You should specify endpoint in start function")
            return

        server = grpc.server(
            futures.ThreadPoolExecutor(max_workers=max_workers),
            options=[('grpc.max_send_message_length', 1024 * 1024 * 1024),
                     ('grpc.max_receive_message_length', 1024 * 1024 * 1024)],
            maximum_concurrent_rpcs=concurrency)
        data_server_pb2_grpc.add_DataServerServicer_to_server(
            DataServerServicer(), server)
        server.add_insecure_port('[::]:{}'.format(endpoint))
        server.start()
        server.wait_for_termination()


if __name__ == "__main__":
    data_server = DataServer()
    data_server.start(endpoint=sys.argv[1])
