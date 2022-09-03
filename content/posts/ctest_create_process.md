---
title: "如何使用 ctest 启动一个服务"
date: 2022-09-03T11:45:03+08:00
draft: false
tags: [cmake]
---

最近在使用 ctest 来创建一个进程，其中过程非常坎坷。这个进程是一个服务，启动后会一直在后台执行，为一些 testcase 交互，比如分配一些资源。

首先发现的 ctest 的一个宏是 `execute_process`，如果是要执行一个普通的命令（比如运行完即可退出的）就非常简单，比如：
```
excute_process(
	COMMAND "ls -l"
	)
```
但是如果要执行一个服务的话，我首先想到的是：
```
execute_process(
	COMMAND "start_process.sh"
)
```
start_process.sh
```
/path/to/the_server_process&
```
实际上这样是不行的，ctest 还是会在 execute_process 处卡住，似乎是在等待被启动进程的退出。
原因是 ctest 在执行进程的时候，会提供自己的 `stdin, stdout, stderr` 给子进程，直到这些 handle 被子进程关闭，才会继续执行后续的指令

```
execute_process(
	COMMAND "start_process.sh"
	INPUT_FILE process.in
  OUTPUT_FILE process.out
  ERROR_FILE process.err
)
```

但是，启动服务这样一种需求不应该使用 `execute_process`来实现，`execute_process`只能在 cmake configure 期间执行，没有合适的机会关闭。 `CTEST_CUSTOM_PRE_TEST `，没有这个问题（可以使用 `POST_TEST` 关闭服务），但是它没有表明这个服务与 test case 之间的关系，并不是所有的 test case 都依赖与该服务。最完美的解决方案是使用 `TEST_FIXTURE`。这个问题断断续续折腾了一个星期，最终因为 stackoverflow 上的一个 comments 第一次提到 `TEST_FIXUTER` 得以完美的解决
