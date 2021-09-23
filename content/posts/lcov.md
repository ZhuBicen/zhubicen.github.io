---
title: "Crash 导致无法生成 gcda 文件"
date: 2021-09-23T22:50:17+08:00
draft: false
tags: [C++]
---

要生成代码的测试覆盖率的 html 报告，一般由以下几个步骤：
1. 使用 `--coverage` 编译源代码
3. 此时会有 .gcno 文件生成
2. 使用 `-lgcov` 来链接生成可执行文件
3. 运行可执行文件，运行相关的测试
4. 检查 .gcda 文件会生成
5. 使用 lcov 生成 info 文件
6. 使用 info 文件生成 html 文件

以上是我积累的处理这个问题的经验。今天刚好有个新的系统要生成自动测试化报告，卡在了第四步，百思不得解

挨个的做了下面的确认：

1. 打开 `ninja -v` 确认 `--coverage` 确认在生效
2. `find . -name "*.gcno"` 也可以看到 gcno 文件生成了

但是运行可执行文件后，就是没有 gcda 文件产生。所以卡在了第四步，试了多遍（重新编译，重新运行）就是不行。反复确认以上的几个步骤都是没问题，一度陷入绝境…… 一直盯着屏幕看 

偶然注意了一个细节，在可执行文件退出的时候 crash 了。会不会是 crash 的时候不会生成 gcda 文件呢，搜索了一下，果然如此。Fix 了退出时 Crash 的问题后，gcda 文件终于生成了。以下的几个步骤就很顺利了

虽然解决这个问题话费了一些时间，大概一个多小时，但是我有了解到了更多的细节，
1. `--coverage` 在编译时相当于选项 `-fprofile-arcs -ftest-coverage`，在链接时相当于 `-lgcov`
2. 在编译对应的 C/CPP 文件时会生成对应的 `gcno` 文件
3. ninja 的 `-v` 选项可以打印出每条编译或者链接的命令
4. 要生成对应的 gcda 文件只需要运行对应的可执行文件就可以了。可以运行多次（执行不同的测试用例），并且可以*同时*运行，gcda 文件可以保证正确的更新。即使对于 Fock 的调用也可以被正确的处理

参考：
[LCOV](http://ltp.sourceforge.net/coverage/lcov.php)
[GCC 文档](https://gcc.gnu.org/onlinedocs/gcc-6.1.0/gcc/Instrumentation-Options.html#Instrumentation-Options)
[Questions on SO](https://stackoverflow.com/questions/10011222/gcov-not-generating-gcda-files)
