---
title: "Compile Gcc4"
date: 2010-01-12T16:58:27+08:00
draft: false
tags: [C++, Linux]
---

最新版本的GCC已经加入了部分C++0X的支持，所以决定自己动手安装最新版GCC

## 准备工作
GCC的编译依赖于GMP和MPRF，如果要编译GMP的话，还要依赖于M4
M4可以这里下载：wget http://ftp.gnu.org/gnu/m4/m4-1.4.10.tar.bz2
gmp和mprf都可以在这里下载：ftp://gcc.gnu.org/pub/gcc/infrastructure
编译：
1. 下载gcc源代码。所有版本的gcc可以在这里下载到ftp://gcc.gnu.org/pub/gcc/releases/
2. tar jxvf gcc-4.4.2.tar.bz2
3. cd gcc-4.4.2
4. ./configure
5. make
6. make install
## 使用
假定新编译的gcc安装在/opt/gcc-4.4.2目录下。
`/opt/gcc-4.42/bin/g++ -std=c++0x -lpthread HelloConcurencyWorld.cpp`
导入新的C++库， export LD_LIBRARY_PATH=$LD_LIBRARY:/opt/gcc-4.4.2/lib
这样就可以运行编译后的C++程序了。
```
void hello()
{
    std::cout << "Hello Concurrent Worldn";
}
int main()
{
    std::thread t(hello);
    t.join();
}
```

虽然步骤相当简单，但是过程中，还是遇到了几个问题:

1. 运行tar jxvf gcc-4.4.2.tar.bz2时，没有对当前目录的权限，导致一大堆的Can not open file错误
2. 运行./configure 时，指定prefix的目录和源码的目录相同了，导致cannot compute suffix of object files: cannot compile。google了半天。
3. 没有指定--enable-language选项，整整编译了一个小时

参考：
http://gcc.gnu.org/install/
http://xahlee.org/UnixResource_dir/_/ldpath.htm