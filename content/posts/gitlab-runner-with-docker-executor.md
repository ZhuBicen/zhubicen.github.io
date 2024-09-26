---
title: "使用 Gitlab Runner Docker Executor"
date: 2024-09-12T19:54:18+08:00
draft: false
tags: []
---

项目的 gitlab pipeline 中有一些是用于测试的，测试在一个 docker container 中进行。目前的脚本是启动 docker container 并且把编译后的二进制和测试代码 mount 到 docker container 中。这样 pipeline 会创建很多的 docker container, 有时候脚本中途退出的话，就会遗留一些 docker container 的进程在。 由于 gitlab 本身提供 docker executor，所以想从 shell executor 转化为 docker executor。转化的过程中还遇到了不少的问题，所以在此记录一下

首先是注册 docker executor gitlab runner，再此步骤中需要提供你的 docker image，注意这个 docker image 只是用来运行你的 gitlab pipeline 脚本的，并不是用来 clone git repo 的 （了解这一点至关重要）。用来 clone 代码的是有另外一个 image 「gitlab-runner-helper」完成。如果你的仓库是用的 ssh 协议，这是就需要 gitlab-runner-helper image 中的 ssh-key 添加到 gitlab 或者 github 中。

以下是用于创建 gitlab-runner-helper image 的 Dockerfile。在编译 Docker Image 过程中，把打印出的 id_rsa.pub 的内容添加到 gitlab/github。

```Dockerfile
FROM registry.gitlab.com/gitlab-org/gitlab-runner/gitlab-runner-helper:x86_64-c1edb478
RUN apk update && apk add openssh-client
# Generate SSH key
RUN ssh-keygen -t rsa -b 4096 -C "systemci@acompany.com" -f /root/.ssh/id_rsa -N ""
# Optional: Display the public key (for debugging purposes)
RUN cat /root/.ssh/id_rsa.pub
RUN ssh-keyscan -H gitlab.acompay.com >> ~/.ssh/known_hosts
```

然后在 gitlab-runner 的配置文件中，指定 gitlab-runner-helper image

```toml
[[runners]]
  name = "huashan"
  url = "https://glhk.acompany.com"
  token = "glrt-_UwHXgVHuUwqit8aNJZ-"
  executor = "docker"
  [runners.custom_build_dir]
  [runners.cache]
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    ######## 这里
    image = "docker-registry/acompany/dev_container:v3"
    helper_image = "docker-registry/acompany/gitlab_runner_helper:latest"
    privileged = true
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
```
