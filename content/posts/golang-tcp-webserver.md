---
title: "一个基于 TCP 实现的 Golang Webserver"
date: 2012-12-02T17:56:17+08:00
draft: false
tags: []
---

这个小程序是从，HTTP: The Definitive Guide,  A Minimal Perl Web Server
移植而来。
``` golang
package main
import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return 
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept err= ", err)
		}
		fmt.Printf("    <<<Request From '%s'>>>\n", conn.RemoteAddr())
		rd := bufio.NewReader(conn)
		
		for {
			line, _, err := rd.ReadLine()
			if err != nil {
				fmt.Println("ReadLine err= ", err)
			}
			if string(line) == "" {
				fmt.Println("")
				fmt.Fprintf(conn, "HELLOWORLD")
				conn.Close()
				break
			}
			fmt.Println(string(line))
		}
	}
}
```
一个`HTTP Request`由`Startline`, `Headers`和`Body`组成。
Startline和Header都是由ASCII字符组成，以行来分隔。每一行都以回车（ASCII 13)和换行(ASCII 10)组成。通常称作*CRLF*。
Body不Startline和Header不同，可以包括字符或者二进制数据，或者为空。

HTTP服务器，就是一个监听指定端口上的TCP连接的程序，检测到连接后，打印出连接端的IP，
然后回显HelloWorld，然后关闭连接。

编译运行该程序后，使用chrome访问，http://192.168.1.3:8080 (192.168.1.3是本机IP）
该服务器有如下打印，可以看到Chrome建立了两个TCP请求，其中一个用来请求favicon.ico
```
    <<<Request From '192.168.1.3:57770'>>>
GET / HTTP/1.1
Host: 192.168.1.3:8080
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip,deflate,sdch
Accept-Language: zh-CN,zh;q=0.8
Accept-Charset: GBK,utf-8;q=0.7,*;q=0.3
```

```
    <<<Request From '192.168.1.3:57771'>>>
GET /favicon.ico HTTP/1.1
Host: 192.168.1.3:8080
Connection: keep-alive
Accept: */*
User-Agent: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11
Accept-Encoding: gzip,deflate,sdch
Accept-Language: zh-CN,zh;q=0.8
Accept-Charset: GBK,utf-8;q=0.7,*;q=0.3

```

