---
title: "Cygwin Gdb"
date: 2011-03-03T00:07:42+08:00
draft: false
tags: [Linux,C++]
---

使用 ygwin 译出来的  序，无法运行，在命令行敲下文件名下，悄无声息的程序就反回了郁闷。加上`g`项，使用gdb加载，出现这个错误：

`db: unknown target exception 0xcxxxxxxx at 0xxxxxxxxxx.`

先查看了程序依赖的dll
```shell
$ cygcheck.exe ./a.exe
D:/tcplex/libssh2/a.exe
  C:/cygwin/bin/cygwin1.dll
    C:/WINDOWS/system32/ADVAPI32.DLL    
      C:/WINDOWS/system32/KERNEL32.dll    
        C:/WINDOWS/system32/ntdll.dll    
      C:/WINDOWS/system32/RPCRT4.dll    
        C:/WINDOWS/system32/Secur32.dll
  C:/cygwin/bin/cygssh2-1.dll
    C:/cygwin/bin/cyggcc_s-1.dll    
    C:/cygwin/bin/cygcrypto-0.9.8.dll    
      C:/cygwin/bin/cygz.dll
  C:/cygwin/bin/cygstdc++-6.dll

> gdb ./a.exe
(gdb) dll cygwin1.dll
(gdb) break *&'dll::init'
Breakpoint 1 at 0x61010270
(gdb) run
```
仍然是提示：`nknown target exception`

 

所以怀疑是cygwin1.dll出了问题。回想最近刚使用新版的setup.exe更新过cygwin，是不是更新程序出了问题？

重新运行cygwin setup.exe把base-cygwin reinstall就可以了。问题解决！
