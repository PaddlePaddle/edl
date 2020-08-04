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
from threading import Lock, Thread
import time
from utils import logger
from paddle_edl.utils.discovery.etcd_client import EtcdClient
import json


class Register(object):
    def __init__(self, etcd_endpoints, job_id, service, server, value):
        self._service = service
        self._server = server
        self._stop = threading.Event()
        self._etcd = EtcdClient(etcd_endpoints, root=job_id, ttl=10)

        if not self._etcd.set_server_not_exists(service_name, server):
            raise exception.EdlRegisterError()

        self._t_register = Threading(self._refresher)

    def _refresher(self):
        while not self._stop.is_set():
            self._etcd_lock.refresh(service_name, server)
            time.sleep(3)

    def stop(self):
        self._stop.set()
        self._t_register.join()


class PodRegister(object):
    def __init__(self, job_env, pod):
        info = self._generate_info(etcd_endpoints, job_env, pod)

        sefl._service_name = "pod"
        self._server = pod._id

        self._register = Register(
            etcd_endpoints=job_env.etcd_endpoints,
            job_id=job_env.job_id,
            service=self._service_name,
            server=self._server,
            info=pod.to_json())

    def stop(self):
        self._register.stop()

    def complete(self):
        info = self._generate_info(
            etcd_endpoints, job_id, pod_id, endpoint, gpus, complete=1)
        self._etcd.set_server_permanent(self._server_name, self._server, info)
        self.stop()


"""
class DataServerRegister(object):
    def __init__(self, etcd_endpoints, job_id, affinity_pod_id,
                 affinity_rank_of_pod, endpoint):
        service_name = "DataServer"
        info = {
            "job_id": str(job_id),
            "affinity_pod_id": str(affinity_pod_id),
            "affinity_rank_of_pod": str(affinity_rank_of_pod)
        }

        server = endpoint

        self._register = Register(
            etcd_endpoints,
            job_id=job_id,
            service=service_name,
            server=server,
            info=json.dumps(info))

    def stop(self):
        self._register.stop()
"""
