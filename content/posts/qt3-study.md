---
title: "Qt3 学习笔记"
date: 2009-12-07T10:56:09+08:00
draft: false
tags: [Linux, C++]
---

QT4已经是如今QT的主流版本，但是仍然有一些老的项目使用QT3。下面记录一下QT3的编译过程：
## 下载源码：
官方提供的QT3源码可以在这里下载到：ftp://ftp.trolltech.no/qt/source/qt-x11-free-3.3.8b.tar.gz
但是其中并没有指明如何编译。SourceForge上有个Q../Free项目，提供了以下下载，方便编译
http://sourceforge.net/projects/qtwin/files/Unofficial%20Qtwin/qt-win-3.3.x-8/qt-3.3.x-p8.7z/download
## 编译QT3
用MingW编译QT3是一个漫长的过程。详细请参考这里：
http://qtwin.sourceforge.net/qt3-win32/compile-mingw.php

Making an Application Debuggable
The release version of an application doesn't contain any debugging symbols or other debugging information. During development it is useful to produce a debugging version of the application that has the relevant information. This is easily achieved by adding debug to the CONFIG variable in the project file.

For example:
```
CONFIG += qt debug 
HEADERS += hello.h 
SOURCES += hello.cpp 
SOURCES += main.cpp
```

Use qmake as before to generate a Makefile and you will be able to obtain useful information about your application when running it in a debugging environment.