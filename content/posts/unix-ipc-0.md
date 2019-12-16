---
title: "Unix Ipc 0"
date: 2009-09-03T13:51:47+08:00
draft: false
tags: [Linux]
---

两种UNIX标准：

1. POSIX，是由IEEE(Institute for Electrical and Electronic Engineers)制定的标准，现由ISO/IEC维护。Posix标准包括3个部分：Part1 和Part2分别被称为Posix.1和Posix.2 System API for C Language Shell and Utility System Administration (正在制定中。。。)
2. X / O p e n是一个国际计算机制造商组织。它提出了一个7卷本可移植性指南X/Open Portability Guide （X / O p e n可移植性指南）第3版〔X／Open 1989〕，我们将其称之为X P G 3。后来又发布了"X/Open Single Unix Specification"后来被称称为Unix95。1997看又发布了Unix98。

另外要指出的是关于 System V， 曾经也被称为 AT&T System V，是Unix操作系统众多版本中的一支，另一支是BSD。它最初由 AT&T 开发，在1983年第一次发布。一共发行了4个 System V 的主要版本：版本1、2、3 和 4。System V Release 4，或者称为SVR4，是最成功的版本，成为一些UNIX共同特性的源头。System v的ipc是open group的单一unix规格版本2[unix 98]所要求的。因此在Unix进程间通信一书中，分别讲述了Posix IPC和System V IPC。

Posix和System V分别支持不同的MessageQueue, Shared Memory和Semaphores.
PIPEs，FIFO,Record Locking机制即不属于Posix标准，也不属于XPG标准。

## 消息传递	
- MessageQueues:	包括Posix MessageQueue,和SystemV MessageQueue
- PIPE, FIFO

## 同步机制	
- Mutex , Condition Variables	
- Read-Write Locks	
- Record Locking	
- Semaphores	
- 共享内存:	包括Posix Semaphores，SystemV Semaphores
- Shared Memory:	包括Posix ShareMemory和SystemVMemory
## 其它	
Doors和SunRPC	



Linux IPC:

最初的Unix发展为两个分支，一个是System V(system IPC)，一个是BSD (socket IPC)。而Linux综了两者的IPC机制。并实现了Posix IPC。
另外请参考：深刻理解Linux进程间通信（IPC)