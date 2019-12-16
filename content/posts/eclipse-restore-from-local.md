---
title: "使用 Eclipse 的 Restore from local history 功能恢复误删除的文件"
date: 2018-03-01T16:26:41+08:00
draft: false
tags: [Linux]
---

在软件开发的过程中，如果不小是删除了正在开发的文件，是个令个苦恼的事情

在文件还没有添加到版本控制系统之前，使用一些命令，比如 `git clean` 删除了不在版本控制系统中的文件，这样会造成文件丢失。由于 Linux 文件系统本身的限制，rm 的文件是很难恢复的。

但是如果你是使用 Eclipse 开发的，有一个`后悔药`可以使用：

- Select Project
- Right Click
- Select Restore from local history in the context menu
- Select your files
- Click OK.

这样误删除的文件就可以恢复了



另外使用 

`git clean -fdx`
之前，可以使用 dry run ,以确认哪些文件会被清理

`git clean -fdxn`