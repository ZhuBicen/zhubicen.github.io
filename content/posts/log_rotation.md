---
title: "Linux logrotation"
date: 2024-09-26T20:53:36+08:00
draft: false
tags: []
---

Linux 上的日志轮转是用 logrotation 实现，logrotation 需要定时任务，根据配置文件定期执行，比如如下的配置文件：

```conf
/var/log/nginx/access.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 644 root root
    sharedscripts
    postrotate
        /usr/local/nginx/sbin/nginx -s reload
    endscript
}
```

注意 logrotation 的执行需要 cron job 或者别的定时器来执行，logrotation 只负责重命名或者压缩文件，重新创建新的 log 文件，需要文件的打印者，比如 nginx 或者 rsyslogd 来重新创建。对于 rsyslogd 可以通过在 postrotate 中发送 HUP 信号给 rsyslogd

logrotation 需要定时执行，所以短时间内如果出现大量的 log，并不能保证 log 很快的 rotation。rsyslogd 可以在这种场景下很好的弥补，比如如下设置，可以在 log 文件达到 20M 的时候，调用 rotate_my_log 轮转

```conf
template(name="CustomFormat" type="string"
  string="%timereported:::date-rfc3339% %syslogtag% [%syslogseverity-text%]\t%msg%\n")

# 20M = 20971520

$outchannel my_log, /var/log/my.log, 20971520, /usr/bin/rotate_my_log

:programname, isequal, "pico_log_forwarder" :omfile:$cluster_log;CustomFormat
& stop
```

rotate_my_log 只需要 `mv /var/log/my.log /var/log/my.log.backup`，这样就可以完成日志在 my.log 和 my.log.backup 文件之间轮转。这种设置在系统磁盘空间受限的时候特别有用。
