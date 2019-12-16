---
title: "在 Docker 中开发 SpringBoot"
date: 2019-08-03T20:46:59+08:00
draft: true
tags: []
---

## 问题

在开发 SpringBoot 微服务的过程中，经常发生的情况是，需要使用 Windows 进行开发，但是微服务最终会部署在 Linux 上，因此而产生的的开发环境和生产环境的不一致，会导致上线后的一些问题


## 解决办法

通过把源代码目录挂载到 Docker 中，并在 Doker 中，运行 `mvn spring-boot:run`，因为 Docker 和 Windows 共享相同的目录，所以 Docker 中的 maven 可以检测到代码的改动，自动重启 SpringBoot 服务

- 假设你的 Spring 代码可以使用 `mvn spring-boot:run` 正常运行。为了使 Maven 可以检测到代码的改动以重启服务，请确保 `pom.xml` 里添加了依赖项：`spring-boot-devtools`
- 假设的你的代码目录在 `C:\Project\springProject` 中。请把下面命令中的目录自行替换为自己项目的目录
- 新建 Docker 目录，并在目录下新建文件 `Dockerfile`

```Dockerfile
FROM maven:3.5.3-jdk-8-alpine
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline
VOLUME /app/src
CMD ["mvn", "spring-boot:run"]
```
- 进入 Docker 目录，编译生成 Docker Image `docker build -t app:dev .`
- 运行 Docker: `docker run -v d:/Project/demo/src:/app/src app:dev`

## 参考
1. https://medium.com/@nieldw/caching-maven-dependencies-in-a-docker-build-dca6ca7ad612
2. https://hub.docker.com/_/maven