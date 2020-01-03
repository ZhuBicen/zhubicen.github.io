---
title: "Docker 的代理设置"
date: 2020-01-03T13:15:41+08:00
draft: false
tags: [Docker]
---

docker 的代理有两部分，一个是docker client，一个是docker daemon：

这里是 `docker daemon` 的设置：https://docs.docker.com/config/daemon/systemd/。 这里的配置主要是`docker pull` 时的网络访问。 主要是设置两个文件：


1. `/etc/systemd/system/docker.service.d/http-proxy.conf`
```
[Service]
Environment="HTTP_PROXY=http://proxy.example.com:80/"
```
2. `/etc/systemd/system/docker.service.d/https-proxy.conf` 
```
[Service]
Environment="HTTPS_PROXY=https://proxy.example.com:443/"
```

3. `sudo systemctl daemon-reload`

4. `sudo systemctl restart docker`

5. `systemctl show --property=Environment docker`

--------------------------------------------------------------------
这里是 `docker client` 的设置，主要是用于 `docker container` 中的网络访问
在 `~/.docker/config.json` 文件中：

```json
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://127.0.0.1:3001",
     "httpsProxy": "http://127.0.0.1:3001",
     "noProxy": "*.test.example.com,.example2.com"
   }
 }
}
```