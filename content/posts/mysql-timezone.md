---
title: "MySQL Timestamp 的时区问题"
date: 2020-03-13T16:09:17+08:00
draft: false
tags: [Java, MySQL]
---

把项目的 Spring Boot 从1.3.5 升级到了 2.2.5，发现一个问题，从数据库里读取的时间都变了。直接使用数据库客户端读取出来的是对的，但是 Java JDBC 读出来的，总是差了几个小时：


为了简单复现问题，用了一小段代码复现，问题：

```Java
jdbcTemplate.queryForObject("SELECT * FROM run WHERE run.id=1", null, (rs, row) -> {
			System.out.println(rs.getTimestamp("startTime"));
			return "";
});
```

直接使用数据库客户端（heidiSQL）读取出来的是： 2020-03-13 15:03:08， 但是 Java 代码读取出来的是： 2020-03-14 04:03:08，相差了 13 个小时

直接使用 SQL ： `SELECT startTime, UNIX_TIMESTAMP(startTime) FROM run WHERE run.sid=6034742;`， 查看到的 UNIX_TIMESTAMP 是： 1584082988， 其对应的时间 是：2020-03-13 07:03:08，转化为 GMT+8 即为：2020-03-13 15:03:08， 这也就是我说的数据库客户端读取出来的时间。 由此看来，MYSQL 写如数据库的时间(timestamp)是按 UTC 写入的，[文档]( https://dev.mysql.com/doc/refman/8.0/en/datetime.html ) 确实是这么写的

> MySQL converts TIMESTAMP values from the current time zone to UTC for storage, and back from UTC to the current time zone for retrieval. (This does not occur for other types such as DATETIME.)

尝试使用如下代码读取，传入时区，就可以读取到和数据库客户端一样的时间：

```
Calendar nowGMTPlus8 = Calendar.getInstance(TimeZone.getTimeZone("GMT+8"));
rs.getTimestamp("startTime"， nowGMTPlus8);
```

到这里就基本确定和时区有关，新的 Java Connector 或许有什么变化，增加了  `serverTimezone=GMT+8` 后问题得以解决

`jdbc:mysql://x.x.x.x:3306/dbname?autoReconnect=true&useSSL=false?&serverTimezone=GMT%2B8`



ServerTimeZone 的说明如下，我疑惑的是为什么之前的版本不需要 ServerTimeZone，新的版本有需要了，服务器的配置没有做任何改动

```
serverTimezone

Override detection/mapping of time zone. Used when time zone from server doesn't map to Java time zone

Since version: 3.0.2
```

https://segmentfault.com/a/1190000016426048

