---
title: "MongoDB 的安装与配置"
date: 2020-03-11T13:45:58+08:00
draft: false
tags: [MongoDB]
---

## 安装 Mongo DB

https://docs.mongodb.com/manual/tutorial/install-mongodb-on-red-hat/

创建文件 `/etc/yum.repos.d/mongodb-org-4.2.repo`，并输入一下内容：
```
[mongodb-org-4.2]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/4.2/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.2.asc
```
运行安装命令： `yum install -y mongodb-org` 

## 启动 Mongo DB 服务

```
systemctl start mongod.service    # For CentOS 8/7 
service mongod restart            # For CentOS 6 

```

Configure MongoDB to autostart on system boot.

```
systemctl enable mongod.service    # For CentOS 8/7 
chkconfig mongod on                # For CentOS 6 
```

##  测试 Mongod DB 数据库

```
[root@tecadmin ~]#  mongo

> use mydb;

> db.test.save( { a: 1 } )

> db.test.find()

  { "_id" : ObjectId("54fc2a4c71b56443ced99ba2"), "a" : 1 }

```

至目前为止，mongodb 还没有鉴权保护，你只能在本机使用 mongodb， 也不需要使用输入任何用户名密码。

接下来我们可以启用用户验证，并从本机外访问 mongodb

首先使用如下命令创建用户：
```
db.createUser({ user: 'root', pwd: 'letmein', roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] });
```


## 使用用户名密码连接

把 bindIp 从 `127.0.0.1` 改为 `0.0.0.0` 以启用从本机以外的连接， `security` 部分增加用户名密码的验证

```

# network interfaces
net:
  port: 27017
  bindIp: 0.0.0.0  
security:
  authorization: enabled

```
重启 mongodb 服务 `service mongod restart`

接下来就可以使用如下的命令连接 mongodb 了
```
mongo  -u root -p letmein
```
你也可以从本机以外连接，使用 `--host` 选项指定 IP 即可

```
mongo --host yourip --port yourport -u root -p letmein
```

## 备份与恢复

可以分别使用 `mongodump` 和 `mongorestore` 来备份和恢复mongo 数据
```
mongodump --host 10.109.3.110 --port 8080 -u mongo_user -p 123456 --authenticationDatabase techdebt -d techdebt

mongoretore 
```


## 客户端工具

推荐使用 `robo 3T`