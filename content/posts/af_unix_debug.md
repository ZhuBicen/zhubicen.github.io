---
title: "记一次 Debug unix local socket 的经历"
date: 2022-09-24T11:07:52+08:00
draft: false
---

测试的 case 是用 Python 写的，通过 AF_UNIX socket 连接到 C++ 的 server 模拟真正的 client 进行测试。其中有一个测试用例通过 unix socket 连续发送了两个消息给 C++ 的 server，理论上来讲 server 应该顺序处理两个消息。但是从测试效果来看，server 处理两个消息的顺序反了

比如 Python 顺序发送了两个消息（在发送两个消息的中间没有任何代码，否侧无法复现），分别是 message1 和 message2，但是 server 先处理了 message2，后处理了 message1。by design（其实是项目的大牛说的） server 应该是先处理 message1 然后处理 message2（事实证明并不是这样，也正是这个先入为主的概念让我走了很多的弯路）

Python 脚本是通过类似下面的程序发送的消息，通过 `create_datagram_endpoint` 连接成功后，程序通过保存的 `transport` 变量来发送消息，比如 `transport.sendto(data, self._remote_addr)`

``` Python
class DatagramEndpointProtocol(asyncio.DatagramProtocol):

    def __init__(self, endpoint):
        pass
    def connection_made(self, transport):
        # send data through transport
        pass
    def connection_lost(self, exec):
        pass
    def datagram_received(self, data, addr):
        pass

 proto = DatagramEndpointProtocol(endpoint)
 sock = socket.socket(socket.AF_UNIX, socket.SOCK_DGRAM)
 os.unlink(maas_local_addr)
 sock.bind(maas_local_addr)
 sock.connect(maas_remote_addr)
 sock.setblocking(False)
 await loop.create_datagram_endpoint(lambda: proto, sock=sock)
```

可以看到这个 socket 是基于 `SOCK_DGRAM` 的，首先想到了会不会是通过 unix local socket 传输的包乱序了呢，就像 UDP 一样不保证到达顺序一样

实际上不是这样，在 UNIX [手册](https://man7.org/linux/man-pages/man7/unix.7.html)上有写:
> preserves message boundaries, and delivers messages in the order that they were sent.

事实上我开始怀疑这是Python 的 asyncio 的 bug（事实再次证明，不要轻易怀疑是系统或者某个库的 bug)，还分别用 C++ 和 Python 的 非 asyncio 写了测试程序，当然是不能复现的

只回到 server 代码本身收取 socket 消息的源头，发现 server 在最初收到消息的时候顺序是没有问题的，然后把收到的消息放进了一个 boost lock free 的 message queue。到这个时候，我又开始怀疑是这个 message queue 的问题，写了测试程序（UT）也测不出问题

只能再回到 server 的代码本身，再往后看一步，server 从 message queue 取出来后也是顺序正确的（Sign，我又走了弯路）。消息的流程进入了一个 Handler，开始在 Handler 里打印点 log，在 Handler 的第一行代码（从 message queue 读消息）加 log，顺序也是正确的，在接下来的几行再次打，发现就乱序了。在这里其实卡了一段时间，最终想明白其实这个时间点，有两个线程在同时调用 Handler 函数（通过加入打印 PID 的代码)，这样就解释的通了

当消息从 local socket 收到，放入 message queue 后，代码会post 消息给 boost::asio 的 io_context，trigger 调用handler 函数，问题就出在这个 io_context 是对应两个线程的。因此会用两个线程同时调用 Handler 函数

在处理这个问题的时候还是遇到了一些挑战的，比如不能在任意的位置随便打印 log，如果在关键的时刻打印 log，就会导致线程的执行不会同时进行，就复现不出问题了。技巧便是创建一个大的数组，简单的把 log 信息放进数组，然后在合适的时候把数组 dump 出来




