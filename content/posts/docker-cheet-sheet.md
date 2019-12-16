---
title: "Docker Cheet Sheet"
date: 2019-08-24T21:06:20+08:00
draft: true
tags: [Docker]
---

# Docker


## Docker commands


### Running a container image


+ `docker run --name kubia-container -p 8080:8080 -d kubia`


    Run a container with bash


+ `docker run --rm -ti ubuntu:latest /bin/bash`


    Remove if exist


+ `docker run -ti --restart=always`


    Restart when failure


To kill a container, run `docker kill xxxcontainer-id`


Remove all containers, run `docker rm $(docker ps -a -q)`


Remove all images, run `docker rmi $(docker images -q -)`








+ `-d` means detach
+ `--name` is the name of the container
+ `kubia` is the image name


### Stop a container


`docker stop kubia-container`


### Start a container


`docker start kubia-container`


## Remove a container


`docker rm kubia-container`


### List all the containers


+ `docker ps -a` show all the docker processes
+ `docker ps` show only the started processes


### Delete all the stopped containers

before performing the removal command, you can get a list of all non-running (stopped) containers that will be removed using the following command:

`docker container ls -a --filter status=exited --filter status=created`

To remove all stopped containers use the docker container prune command:

`docker container prune`

## How to push a docker image


```bash
docker login
docker tag firstimage YOUR_DOCKERHUB_NAME/firstimage
docker push YOUR_DOCKERHUB_NAME/firstimage
```




## Minikube


`minikube start --docker-env http_proxy=http://127.0.0.1:1080 --docker-env https_proxy=http://127.0.0.1:1080`


`minikube start --docker-env http_proxy=http://127.0.0.1:1080 --docker-env https_proxy=https://127.0.0.1:1080`


[Minikube behind http proxy](https://codefarm.me/2018/08/09/http-proxy-docker-minikube/)


```
minikube start --docker-env HTTP_PROXY=http://127.0.0.1:1080/ --docker-env HTTPS_PROXY=http://127.0.0.1:1080/ --docker-env NO_PROXY=index.docker.io,registry.hub.docker.com,
 registry-1.docker.io,registry.docker-cn.com,
 registry-mirror-cache-cn.oss-cn-shanghai.aliyuncs.com,
 192.168.99.100 --registry-mirror https://registry.docker-cn.com
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbODc5NjE5NjldfQ==
-->
