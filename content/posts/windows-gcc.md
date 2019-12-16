---
title: "Windows Gcc"
date: 2009-06-06T08:31:40+08:00
draft: false
tags: [C++]
---

Visual Studio 变得越来越大，其中附带了太多的与标准C++无关的东西，有时只是想些一个小的程序，却要启动一个又笨又重的IDE，实在是不方便。还有安装VS就要花相当的时间和相当的磁盘空间。

除了VS我们还有很多其它的选择：MinGW，Cygwin，Dev cpp，CodeBlocks，MinW Studio。其中有一些是IDE。基于以下的原因，我选择了MinGW：

1． 如侯捷所言，“在 console mode 中使用 C/C++ 编译器”

2． MinGW是著名编译器GCC的Windows移植，所以所有在該编译器下的经验都可以，应用到Linux平台。于我而言，我已习惯了Linux下使用Makefile控制程序的编译过程，所以这些经验可以在Windows平台下得以使用。

3． MinW中的MSYS可以很好的建立起类似Linux的Shell，这样几乎所有的Linux命令可以在Windows下得以应用。

4． 和Cygwin不同，使用MinW编译后的可执行文件，不需要依赖于任何库文件就可以在Windows平台运行。而Cygwin需要依赖于Cygwin.dll

接下来，让我们开始MinW的安装吧。这里是MinW的Getting Started。如果你E文不怎么好，请看以下的安装步骤：

1． 在这里下载安装文件并运行，安装过程中有以下几个组件可供选择安装：

A．MinGW base tools：这里包括MinGW的一些基础的二进制程序，和C编译器，以及平台API

B． …Compiler，如果只要编译C++，就选g++编译器

C． MinGW Make 是linux下Make的移植，建议选择

提醒不要把安装路径选择到有空格的路径。

2． 安装MSYS，点击此处下载。在安装的最后步骤将会执行MSYS的post install script。在这里请选择你所安装的MSYS的目录，这样mingw将会被挂载到/mingw。我们也就可以MSYS中使用gcc编译器了。如不清楚，请阅读安装后的ReadMe。

打开MSYS，我们的编程的工作就可以在此展开了。
