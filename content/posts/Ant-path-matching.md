---
title: "Ant Path Matching"
date: 2020-01-16T16:51:07+08:00
draft: true
tags: [Spring Cloud]
---

在使用 Spring Cloud Gateway 的[Path Route Predicate](https://cloud.spring.io/spring-cloud-gateway/reference/html/#the-path-route-predicate-factory) 的时候，发现了一个现象，`/content/**` 可以匹配 到`/content`  

按照之前我的理解，`/content**` 才可以匹配`/content`。之后发现这是 `Ant Path Matcher` 的 定义，`＊*` 可以匹配 0 到多个目录

