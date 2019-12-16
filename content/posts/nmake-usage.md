---
title: "Nmake Usage"
date: 2008-12-16T20:11:21+08:00
draft: false
tags: [C++]
---

现在维护一个旧的支持多国语言的程序（使用vc6.0)，由于程序比较早，采用的方法也比较原始，因此每次要build一个语言的版本，就要改变一下编译的参数。比如，编译一个中文版的时候就要预定义：LAN_CHS，要编译韩文版的时候就要定义：LAN_KOR。这样来回的折腾实在是太麻烦了，我因此导出了该项目的makefile，通过增加一个参数的方式，可以指定要编译的语言版本，因此：nmake LAN=”CHS”就得到了一个中文版，nmake LAN=”KOR”就得到了一个韩文版：）

为什么现在的vc2005, 2008不支持导出makefile了呢。当然我们有msbuild, devenv可以使用。

nmake的参考在这里：http://msdn.microsoft.com/en-us/library/dd9y37ha(VS.80).aspx
