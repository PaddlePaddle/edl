# PaddlePaddle Elastic Deep Learning

<img src="../logo/edl.png" width="500">

PaddlePaddle (PArallel Distributed Deep LEarning) is an easy-to-use, efficient,
flexible and scalable deep learning platform, which is originally developed by
Baidu scientists and engineers for the purpose of applying deep learning to many
products at Baidu.

PaddlePaddle Elastic Deep Learning (EDL) is a clustering project which leverages PaddlePaddle training jobs to
be scalable and fault-tolerant. EDL will greatly boost the parallel distributed training jobs and make good use
of cluster computing power.

EDL is based on the full fault-tolerant feature of PaddlePaddle, it uses a Kubernetes controller to manage
the cluster training jobs and an auto-scaler to scale the job's computing resources.

For researchers, EDL with Kuberntes will reduce the waiting time of the job submitted, to help with
exposing potential algorithmic problems as early as possible.

For enterprises, a complete data pipeline includes training jobs, web servers,
log collector and so on. These components often run on a distributed operation system
like k8s. EDL make it possible to run less deep learning job processes during
periods of high web traffic, more when web traffic is low. EDL would optimize the global
utilization of a cluster.

## Tutorial Outline

- Introduction
  At the introduction session, we will introduce:
    - The latest PaddlePaddle version Fluid; and
    - Why we develop PaddlePaddle EDL and how we implement it.
- Hands-on Tutorial
  Following the introduction, we have a hands-on tutorial after each introduction
  session so that all the audience can use PaddlePaddle and ask some questions
  while using PaddlePaddle:
    - Training models using PaddlePaddle Fluid.
    - Launch EDL training jobs on a Kubernetes cluster.

## Prerequisites

- [Install Docker](https://docs.docker.com/install/)
- [Install kubectl](./install.md#kubectl)
- [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- A Kubernetes cluster which version is `1.7.x`:
    - [minikube would launch a kubernetes cluster locally](./install.md#minikube).
    - [kops would launch a Kuberntes cluster on AWS](./install.md#aws).
    - We will also prepare a public Kubernetes cluster via Cloud if you don't have an AWS
      account that you can submit the EDL training jobs using the public cluster.

## Resources

- [PaddlePaddle](http://github.com/PaddlePaddle/Paddle)
- [PaddlePaddle Book](http://github.com/PaddlePaddle/book)
- [PaddlePaddle EDL](https://github.com/PaddlePaddle/edl)

## Part-1 Training Models on Your Laptop using PaddlePaddle

Please checkout [PaddlePaddle Book](http://github.com/PaddlePaddle/book), steps to run
the training process and example output.

## Part-2: Launch the PaddlePaddle EDL Training Jobs on a Kubernetes Cluster

Please note, EDL only support the early PaddlePaddle version so the fault-tolerant model is
written with PaddlePaddle v2 API.

### Configure kubectl

If you start up a Kubernetes instance by `minikube` or `kops`, the kubectl configuration would be ready when
the cluster is available, for the other approach, you can contact the administrator to fetch the configuration file.

### Deploy EDL Components

**NOTE**: there is only one EDL controller in a Kubernetes cluster, so if you're using a public cluster, you can skip this step.

1. (Optional) Configure RBAC for EDL controller so that it would have the cluster admin permission.

    If you launch a Kubernetes cluster by kops on AWS, the default authenticating policy is `RBAC`, so this step is **necessary**:

    ```bash
    kubectl create -f k8s/rbac_admin.yaml
    ```

1. Create TPR "Training-Job"

    ``` bash
    kubectl create -f k8s/thirdpartyresource.yaml
    ```

    To verify the creation of the resource, run the following command:

    ``` bash
    kubectl describe ThirdPartyResource training-job
    ```

- Deploy EDL controller

    ```bash
    kubectl create -f k8s/edl_controller.yaml
    ```

### Launch the EDL Training Jobs

1. Edit the local training program to be able to run with distributed mode

    It's easy to update your local training program to be running with distributing mode:

    - Dataset

        Pre-process the dataset to convert to RecordIO format,
        We have done this in the Docker image `paddlepaddle/edl-example` using `dataset.covert` API as follows:

        ``` python
        dataset.common.convert('/data/recordio/imikolov/', dataset.imikolov.train(word_dict, 5), 5000, 'imikolov-train')"
        ```

        This would generate many recordio files on `/data/recordio/imikolov` folder, and we have prepared these files on Docker image `paddlepaddle/edl-example`.

    - Pass the `etcd_endpoint` to the `Trainer` object so that `Trainer` would know it's a fault-tolerant distributed training job.

        ``` python
        trainer = paddle.trainer.SGD(cost,
                                      parameters,
                                      adam_optimizer,
                                      is_local=False,
                                      pserver_spec=etcd_endpoint,
                                      use_etcd=True)
        ```

    - Use `cloud_reader` which is a `master_client` instance can fetch the training data from the task queue.

        ``` python
        trainer.train(
            paddle.batch(cloud_reader([TRAIN_FILES_PATH], etcd_endpoint), 32),
            num_passes=30,
            event_handler=event_handler)
        ```

1. Run the monitor program

    Please open a new tab in your terminal program and run the monitor Python script `example/collector.py`:

    ```bash
    docker run --rm -it -v $HOME/.kube/config:/root/.kube/config $PWD:/work paddlepaddle/edl-example python collector.py
    ```

    And you can see the following metrics:

    ``` text
    SUBMITTED-JOBS    PENDING-JOBS    RUNNING-TRAINERS    CPU-UTILS
    0    0    -    18.40%
    0    0    -    18.40%
    0    0    -    18.40%
    ...
    ```

1. Deploy EDL Training Jobs

    ```bash
    kubectl create -f example/examplejob.yaml
    ```

1. Deploy Multiple Training Jobs and Check the Monitor Logs

    You can edit the YAML file and change the `name` field so that you can submit multiple training jobs.
    For example, I submited three jobs which name is `example`, `example1` and `example2`, the monitor logs
    is as follows:

    ``` text
    SUBMITED-JOBS    PENDING-JOBS    RUNNING-TRAINERS    CPU-UTILS
    0    0    -    18.40%
    0    0    -    18.40%
    1    1    example:0    23.40%
    1    0    example:10    54.40%
    1    0    example:10    54.40%
    2    0    example:10|example1:5    80.40%
    2    0    example:10|example1:8    86.40%
    2    0    example:10|example1:8    86.40%
    2    0    example:10|example1:8    86.40%
    2    0    example:10|example1:8    86.40%
    3    1    example2:0|example:10|example1:8    86.40%
    3    1    example2:0|example:10|example1:8    86.40%
    3    1    example2:0|example:5|example1:4    68.40%
    3    1    example2:0|example:3|example1:4    68.40%
    3    0    example2:4|example:3|example1:4    88.40%
    3    0    example2:4|example:3|example1:4    88.40%
    ```

- At the begging, then there is no training job in the cluster except some Kubernetes system components, so the CPU utilization is **18.40%**.
- After submitting the training job **example**, the CPU utilization rise to **54.40%**, because of the `max-instances` in the YAML file is **10**, so the running trainers is **10**.
- After submitting the training job **example1**, the CPU utilization rose to **86.40%**.
- While we submitting the training job **example2**, there is no more resource for it, so EDL auto-scaller would
scale down the other jobs' trainer process, and eventually the running trainers of **example** dropped down to **3**, **example1** dropped down to **4** and no pending jobs.