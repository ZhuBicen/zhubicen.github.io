---
title: "使用 Gitlab Runner With Docker Executor"
date: 2024-09-12T19:54:18+08:00
draft: true
tags: []
---

项目的 gitlab pipeline 中有一些是用于测试的，测试在一个 docker container 中进行。目前的脚本是启动 docker container 并且把编译后的二进制和测试代码 mount 到 docker container 中。这样 pipeline 会创建很多的 docker container, 有时候脚本中途退出的话，就会遗留一些 docker container 的进程在。 由于 gitlab 本身提供 docker executor，所以想从 shell executor 转化为 docker executor。转化的过程中还遇到了不少的问题，所以在此记录一下

首先是注册 docker executor gitlab runner，再次步骤中需要提供你的 docker image，注意这个 docker image 只是用来运行你的 gitlab pipeline 脚本的，并不是用来 clone git repo 的，这一点至关重要。
