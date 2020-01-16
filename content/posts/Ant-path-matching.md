---
title: "Ant Path Matching"
date: 2020-01-16T16:51:07+08:00
draft: true
tags: []
---

在使用 Spring Cloud Gateway 的[[Path Route Predicate](https://cloud.spring.io/spring-cloud-gateway/reference/html/#the-path-route-predicate-factory) 的时候，发现了一个现象，`/content/**` 可以匹配 到`/content` 。 之后发现