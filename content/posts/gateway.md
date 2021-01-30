---
title: "Sprint Cloud Gateway"
date: 2021-01-30T11:00:53+08:00
draft: false
tags: [Spring, Sprint Cloud]
---

最近有这么一个需求，使用 Spring CLoud Gateway 把以 `/exchange/` 开头的请求转发到一个服务器，以 `/reservation/` 开头的请求转发到另外一个服务器，很自然的使用下面的配置文件：

```yml
spring:
   cloud:
    gateway:
      routes:
        - id: exchange
          uri: http://localhost:8002/
          predicates:
            - Path=/exchange/**
        - id: reservation
          uri: http://localhost:8003/
          predicates:
            - Path=/reservation/**
```

这里直接 hardcode 了相关的 URL（包括 IP 和端口），测试没有问题。由于项目中还用了 NACOS 作为服务注册于发现中心，因此不能直接 Hardcode URL，需要使用 Load Balance 的方式。根据自己的印象改成了下面这样：

```yml
spring:
   cloud:
    gateway:
      routes:
        - id: exchange
          uri: lb://localhost:8002/
          predicates:
            - Path=/exchange/**
        - id: reservation
          uri: lb://localhost:8003/
          predicates:
            - Path=/reservation/**
```

上面的配置文件犯了两个错误，一个是没有使用服务名，服务名是注册在服务发现中心的名字，另一个是使用了端口号，gateway 使用服务名获取对应服务的IP和端口，所以这里是不需要端口号的，因此正确的配置文件如下，`exchange` 和 `reservation` 都是服务名：
```yml
spring:
   cloud:
    gateway:
      routes:
        - id: exchange
          uri: lb://exchange
          predicates:
            - Path=/exchange/**
        - id: reservation
          uri: lb://reservation
          predicates:
            - Path=/reservation/**
```
