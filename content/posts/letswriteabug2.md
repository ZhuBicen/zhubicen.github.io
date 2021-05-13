---
title: "来，写个 bug  -- 间接引用坏指针"
date: 2021-05-13T21:53:51+08:00
draft: false
tags: []
---

在进程的虚拟地址空间中有较大的洞，没有映射到任何有意义的数据。如果我们试图引用一个指向这些洞的指针，那么操作系统就会以段异常 (Segmentation Fault) 中止我们的程序。而且，虚拟存储器的某些区域是只读的。试图写这些区域也将会以保护异常中止这个程序。


Segmentation fault 是由于访问了“不属于”自己的内存引起的。这时一种机制用来阻止程序进入一种内存被破坏并且很难调试的状态。当程序遭遇了 Segmentation fault 时，你就知道你的程序访问了已经被释放的内存，或者“只读”的内存

最典型的情况是访问了空指针


```C
int *p = NULL;
*p = 1;
```

另外一种情况是访问了“只读” 的内存

```C
char *str = "Foo"; // Compiler marks the constant string as read-only
*str = 'b'; // Which means this is illegal and results in a segfault
```

还有指向不存在的对象的悬空指针(Dangling Point)，像这样：

```C
char *p = NULL;
{
    char c;
    p = &c;
}
// Now p is dangling
```

指针 p 悬空因为它指向了变量 c ，而 c 在过了作用域后已经被析构了。在出了作用域后如果使用 p （比如: *p = 'A') 也会造成 Segmentation fault


参考：[https://stackoverflow.com/questions/2346806/what-is-a-segmentation-fault](https://stackoverflow.com/questions/2346806/what-is-a-segmentation-fault)