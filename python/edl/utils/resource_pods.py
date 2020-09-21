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
from edl.utils import register

class PodResourceRegister(register.Register):
    def __init__(self, job_env, pod):
        service = constants.ETCD_POD_RESOURCE
        server = "{}".format(pod.get_id())
        value = pod.to_json()

        super(PodResourceRegister, self).__init__(
            etcd_endpoints=job_env.etcd_endpoints,
            job_id=job_env.job_id,
            service=service,
            server=server,
            info=value)

@error_utils.handle_errors_until_timeout
def get_resource_pods_dict(etcd, timeout=15):
    servers = etcd.get_service(constants.ETCD_POD_RESOURCE)

    pods = {}
    for s in servers:
        p = pod.Pod()
        p.from_json(s.info)
        pods[p.get_id()] = p

    return pods

@error_utils.handle_errors_until_timeout
def wait_resource(pod_id, timeout=15):
    pods = get_resource_pods_dict(timeout=timeout)
    if len(pods) == 1:
        if pod_id in pods:
            return True

    if len(pods) == 0:
        return True

    return False