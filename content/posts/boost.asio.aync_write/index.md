---
title: "使用 Boost.asio.async_write 时的小问题"
date: 2021-11-23T09:57:39+08:00
draft: false
tags: [C++]
---

遇到了一个小问题，在使用 `boost::asio::async_write` 的时候，发现在数据量比较大的时候会出现发送的字符数据乱码的问题。在调用真正的发送函数（也就是 `async_write`) 之前，打印出的发送数据都是正常的，似乎是 `async_write` 出了问题，从代码上来看，在发送前数据已经正确的传给了 `boost::asio::buffer`，然后调用就出现了问题。

其实这是对 `boost::asio::buffer` 的理解错误，基于效率的考虑，async_write 并不会把数据拷贝一遍，而只是引用了数据地址和对应的长度


```c++
boost::asio::async_write(socket, 
        boost::asio::buffer(buffer->c_str(), buffer->size()), 
        boost::asio::transfer_at_least(inBuffer->size()), 
        std::bind(&LocalHandle::HandleWrite, shared_from_this(), _1, _2));
```

查看 `async_write` 的文档是这么描述第二个参数的：
```
One or more buffers containing the data to be written. Although the buffers object may be copied as necessary, ownership of the underlying memory blocks is retained by the caller, which must guarantee that they remain valid until the handler is called.
```

而在我们的代码中，buffer 是指向一个智能指针指向的内存。在数据量小的时候即使智能指针已经释放了但是其中的内存还没有被分给别的对象使用，因此碰巧没有出问题。但是在数据量比较大的时候，已经被释放的内存迅速被别的对象占用，因此在真正发送的时候就会使用到这些内存，导致出现乱码

附一张当时使用 socat 工具抓取的 Unix Domain Socket 的数据截图：

![image-20211219100055711](./boost.asio.aync_write.assets/image-20211219100055711.png)

> 另外还遇到一个小问题，关于 `std::vector`，由于最大长度是固定的所以就直接 `std::vector::reserve` 了大小，后面就直接使用了。`vector` 中存储的是 `Lambda Object` ，事先 reserve 的大小是 1000，使用如下的代码赋值，`v[i] = [](){}` 真实的 lambda 还是有点复杂的，crash 时候的堆栈指向了 vector 的 `operator=` 函数，因此怀疑是拷贝 lambda 的时候产生了异常。但是把 lambda 精简到非常简单还是出问题。最终幡然醒悟 reserve 的空间并不能直接使用，因为其内存没有初始化对象。问题的根源是这种场景下应该使用 `std::array` 啥问题没有



