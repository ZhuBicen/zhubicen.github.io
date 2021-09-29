---
title: "asyncio.wait_for 的一个 bug ？"
date: 2021-09-29T22:01:58+08:00
draft: False
tags: [Python]
---

从第一次学习使用 Python 到现在已有十余年了，这些年中写个小脚本，读写个数据库都算是顺手，从来没有遇到过什么坑，直到昨天开始使用 asyncio。在没有使用 asyncio 前，在书上也读到过这个主题，隐约觉得有些复杂，也可能是 goroutine 先入为主的原因，相比较于 goroutine，coroutine 太难理解了，yield，yield from，总是模糊搞不清楚，或许这就为踩坑埋下了伏笔



这次正经使用（以前随便写点小脚本，都算不正经） Python 是为了测试一个 C++ 些的程序，该程序使用 Unix Socket （也叫 Unix Domain Socket，我们项目组的同事都称之为 local socket，现在想来其实不够准确）和在同一宿主机上的一个 go 语言些的程序进行进程间通信。如果你还不知道什么是 Unix Socket，它就是创建 Socket 的时候 Socket Familiy 指定为 `AF_UNIX `的 Socket ，另一个大名鼎鼎的就是 `AF_INET`。因此这个测试程序的任务就是通过 Unix Socket 和被测试程序通信，发送接受消息和数据



由于对 asyncio 没有把握，因此第一个版本就是阻塞IO写的，使用 `socket.read()` `socket.write`随着测试用例的增加，阻塞版本已经无法满足需求，比如需要同时收取消息和发送数据，还要有超时的功能。没办法，只能改成 asyncio，把 `socket.read()` 改成了 `loop.sock_read`，把 `socket.write()` 改成了 `loop.sock_write`。改这些（包括给一些函数加上 async, await），都还算顺利。但是在使用 `asyncio.wait_for`的时候就卡壳了，无论如何它就是不会超时。多次尝试无果后，只能写了个小程序把它复现了（同时也去除了使用的Pytest)，就是下面的程序。

背景：对端的服务器在连接成功后的很久（100秒）才发送数据，但是这边的3秒的超时设置就是不起效，一直会等到对端服务器100秒后收到数据才返回，也就是总没有 timeout

```python
import sys
import asyncio
import socket

async def test_single_call(loop):
    server_address = './unixSocket'
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.connect(server_address)
    try:
        msg = await asyncio.wait_for(loop.sock_recv(sock, 1), timeout=3)
        print("Unexcepted message received:" , msg, file=sys.stderr)
        assert False
    except asyncio.TimeoutError:
        pass
    msg = await loop.sock_recv(sock, 1)

loop = asyncio.get_event_loop()
loop.run_until_complete(test_single_call(loop))
loop.close()
```

这时候读了 `loop.sock_recv` 相关的文档，发现另一个函数 `loop.sock_connect` 代码变成了这样：

```python
import sys
import asyncio
import socket

async def test_single_call(loop):
    server_address = './unixSocket'
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    # sock.connect(server_address)
    # sock.setblocking(False) ?
    sock = await loop.sock_connect(sock, server_address)
    try:
        msg = await asyncio.wait_for(loop.sock_recv(sock, 1), timeout=3)
        print("Unexcepted message received:" , msg, file=sys.stderr)
        assert False
    except asyncio.TimeoutError:
        pass
    msg = await loop.sock_recv(sock, 1)

loop = asyncio.get_event_loop()
loop.run_until_complete(test_single_call(loop))
loop.close()

```

这样竟然就可以了，使用 asyncio 的 tcp 客户端。这是我在写这篇博客的时候发现的。我当时走了很多的弯路，除了试了 raw socket （也就是上面的代码）还试了 `open_connection, open_unix_connection`，open_unix_connection 是可以超时成功的，但是不会像文档里说的那样在超时后自动 cancel task，为了这个问题我还专门在 Stack overflow 上[提问](https://stackoverflow.com/questions/69370252/asyncio-wait-for-didnt-cancel-the-task)，后来更多的尝试证明这是 Python 3.6  (Linux 版) 的一个问题，在 Python 3.8，3.9 上都是可以自动 Cancel 的，总之此处省略一千字...



这就是曲折的过程，我只是需要：超时，并且取消一个 coroutine，每一步都不顺利（版本也是问题，Ubuntu 18.04 上的Python 3.6 太老？），只此记录，以备自省（以后碰到问题先以最小的程序复现，也就离解决不远了）

