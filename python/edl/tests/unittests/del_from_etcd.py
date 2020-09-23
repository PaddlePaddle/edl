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

from edl.utils import constants
from edl.discovery import etcd_client

g_etcd_endpoints = "127.0.0.1:2379"

job_id = os.environ["PADDLE_JOB_ID"]
etcd_endpoints = os.environ["PADDLE_ETCD_ENDPOINTS"]
etcd = etcd_client.EtcdClient([g_etcd_endpoints], root=job_id)
etcd.init()
constants.clean_etcd(etcd)