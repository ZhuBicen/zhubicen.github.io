---
title: "git 常用配置"
date: 2011-12-16T13:43:27+08:00
draft: false
tags: [git]
---

# 代理

```Shell
git config --global https.proxy http://127.0.0.1:1080
git config --global https.proxy http://127.0.0.1:1080
git config --global --unset http.proxy
git config --global --unset https.proxy
```



# 别名

比如,  从 svn 转到 git 的应该深有同感

```Shell
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```


git 常用操作示意图:

![1576475173484](./git-alias.assets/1576475173484.png)