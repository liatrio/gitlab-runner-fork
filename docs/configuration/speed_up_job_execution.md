---
stage: Verify
group: Runner
info: To determine the technical writer assigned to the Stage/Group associated with this page, see https://about.gitlab.com/handbook/engineering/ux/technical-writing/#designated-technical-writers
---

# Speed up job execution

You can improve performance of your jobs by caching your images and dependencies.

## Use a proxy for containers

You can speed up the time it takes to download Docker images by using:

- The GitLab Dependency Proxy or 
- A mirror of the DockerHub Registry

### GitLab Dependency Proxy

To more quickly access container images, you can
[use the Dependency Proxy](https://docs.gitlab.com/ee/user/packages/dependency_proxy/)
to proxy container images.

### Docker Hub Registry mirror

You can also speed up the time it takes for your jobs to access container images by mirroring Docker Hub.
This results in the [Registry as a pull through cache](https://docs.docker.com/registry/recipes/mirror/).
In addition to speeding up job execution, a mirror can make your infrastructure
more resilient to Docker Hub outages and Docker Hub rate limits.

When the Docker daemon is [configured to use the mirror](https://docs.docker.com/registry/recipes/mirror/#configure-the-docker-daemon)
it automatically checks for the image on your running instance of the mirror. If it's not available, it
pulls the image from the public Docker registry and stores it locally before handing it back to you.

The next request for the same image pulls from your local registry.

More detail on how it works can be found [here](https://docs.docker.com/registry/recipes/mirror/#how-does-it-work).

#### Use a Docker Hub Registry mirror

To create a Docker Hub Registry mirror:

1. Log in to a dedicated machine where the proxy container registry will run.
1. Make sure that [Docker Engine](https://docs.docker.com/install/) is installed
   on that machine.
1. Create a new container registry:

   ```shell
   docker run -d -p 6000:5000 \
       -e REGISTRY_PROXY_REMOTEURL=https://registry-1.docker.io \
       --restart always \
       --name registry registry:2
   ```

   You can modify the port number (`6000`) to expose the registry on a
   different port. This will start the server with `http`. If you want
   to turn on TLS (`https`) follow the [official
   documentation](https://docs.docker.com/registry/configuration/#tls).

1. Check the IP address of the server:

   ```shell
   hostname --ip-address
   ```

   You should choose the private network IP address. The private
   network is usually the fastest solution for internal communication
   between machines on a single provider, like DigitalOcean, AWS, or Azure.
   Usually, storage on a private network is not applied against your monthly
   bandwidth limit.

The Docker Hub registry is accessible under `MY_REGISTRY_IP:6000`.

You can now [configure `config.toml`](../configuration/autoscale.md#distributed-container-registry-mirroring)
to use the new registry server.

## Use a distributed cache

You can speed up the time it takes to download language dependencies by
using a distributed [cache](https://docs.gitlab.com/ee/ci/yaml/#cache).

To specify a distributed cache, you set up the cache server and then
[configure runner to use that cache server](../configuration/advanced-configuration.md#the-runnerscache-section).

If you are using autoscaling, learn more about the distributed runners
[cache feature](../configuration/autoscale.md#distributed-runners-caching).

The following cache servers are supported:

- [AWS S3](#use-aws-s3)
- [MinIO](#use-minio) or other S3-compatible cache server
- [Google Cloud Storage](#use-google-cloud-storage)
- [Azure Blob storage](#use-azure-blob-storage)

Learn more about GitLab CI/CD [cache dependencies and best practices](https://docs.gitlab.com/ee/ci/caching/index.html).

### Use AWS S3

To use AWS S3 as a distributed cache,
[edit runner's `config.toml` file](../configuration/advanced-configuration.md#the-runnerscaches3-section) to point
to the S3 location and provide credentials for connecting.

### Use MinIO

Instead of using AWS S3, you can create your own cache storage.

1. Log in to a dedicated machine where the cache server will run.
1. Make sure that [Docker Engine](https://docs.docker.com/install/) is installed
   on that machine.
1. Start [MinIO](https://min.io), a simple S3-compatible server written in Go:

   ```shell
   docker run -it --restart always -p 9005:9000 \
           -v /.minio:/root/.minio -v /export:/export \
           --name minio \
           minio/minio:latest server /export
   ```

   You can modify the port `9005` to expose the cache server on a
   different port.

1. Check the IP address of the server:

   ```shell
   hostname --ip-address
   ```

1. Your cache server will be available at `MY_CACHE_IP:9005`.
1. Create a bucket that will be used by the Runner:

   ```shell
   sudo mkdir /export/runner
   ```

   `runner` is the name of the bucket in that case. If you choose a different
   bucket, then it will be different. All caches will be stored in the
   `/export` directory.

1. Read the Access and Secret Key of MinIO and use it to configure the Runner:

   ```shell
   sudo cat /export/.minio.sys/config/config.json | grep Key
   ```

You can now
[configure `config.toml`](../configuration/autoscale.md#distributed-runners-caching)
to use the new cache server.

### Use Google Cloud Storage

To use Google Cloud Platform as a distributed cache,
[edit runner's `config.toml` file](../configuration/advanced-configuration.md#the-runnerscachegcs-section) to point
to the GCP location and provide credentials for connecting.

### Use Azure Blob storage

To use Azure Blob storage as a distributed cache,
[edit runner's `config.toml` file](../configuration/advanced-configuration.md#the-runnerscacheazure-section) to point
to the Azure location and provide credentials for connecting.