---
title: "Heap Corruption"
date: 2006-11-10T22:58:46+08:00
draft: false
tags: [C++]
---

昨天遇到一个HEAP CORRUPTION错误，花了好多的时间才找到原因，现总结如下，希望大家遇到同样的问题的时候，能迅速定位错误的代码。

错误的现象是这样的：

在程序的开始处我申请了一块内存，中间对其进行了一些操作，在程序结束处，释放内存的时候，引起错误：

```
HEAP CORRUPTION DETECTED：after Normal block(#***) at 0x****.CRT detected that application wrote memory after end of heap buffer.
```

错误原因：

以对内存操作的过程中，所写的地址超出了，所分配内存的边界

解决办法：

在可能出错的代码处，使用_CrtCheckMemory进行检测

比如：

```
 int* p = new int[2];
 *(p+2) = 1;
 _ASSERTE( _CrtCheckMemory( ) );
```