---
title: "Ubuntu Docker Timezone"
date: 2019-07-26T10:45:18+08:00
draft: false
tags: [Docker]
---

一般来讲如下方法就够了：
```
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata
```

但是因为 Ubuntu 的 bug,  要添加一行 `RUn rm -f /etc/localtime`, 因此:

```
RUN echo "Asia/Shanghai" > /etc/timezone
RUN rm -f /etc/localtime    
RUN dpkg-reconfigure -f noninteractive tzdata
```

以上参考自：[docker-timezone-in-ubuntu-16-04-image](https://stackoverflow.com/questions/40234847/docker-timezone-in-ubuntu-16-04-image), 虽然原文只是针对 1604, 笔者亲测1804也是适用的