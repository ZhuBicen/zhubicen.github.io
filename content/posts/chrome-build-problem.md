---
title: "Chrome Build Problem"
date: 2018-04-05T10:30:44+08:00
draft: false
tags: [C++,Chrome]
---

下载代码前最基本的代理设置：https://blog.csdn.net/siyu77/article/details/50916320
对于 ShadowSocks 代理 https_proxy 也要设置成  http://localhost:1080（注意代理的 url 没有s)
gclient sync 失败的问题(Failed to fetch file gs://chromium-gn/a68b194afb05d6a6357cf2e2464136ed7723c305) 如下图，还需要设置 boto 代理

建立 .poto 文件如下（文件名不重要）
```
[Boto] 
 proxy = localhost 
 proxy_port = 1080 
```
然后使用环境变量 NO_AUTH_BOTO_CONFIG 指向 boto 文件的位置
`export NO_AUTH_BOTO_CONFIG=C:/path/to/.boto` 
来源： https://groups.google.com/a/chromium.org/forum/#!topic/chromium-packagers/0eE3oC3lBFw

另外还有一个 Python Unicode 的问题：https://groups.google.com/a/chromium.org/forum/#!topic/chromium-dev/daWDn2cfSyA
目前（2018/4/5)还是需要手工的合并 Fix:
可能英文版的没有这个问题，所以这个 bug 一直没有合并

另外我还遇到了执行 `“gn gen --ide=vs out\Default”` 出错的问题，原因是因为没有安装 10.0.15063 版本的 SDK.

> \"/Ic:\\program files (x86)\\microsoft visual studio\\2017\\community\\vc\\tools\\msvc\\14.13.26128\\include\" 
\"/Ic:\\program files (x86)\\windows kits\\10\\include\\10.0.16299.0\\ucrt\" 
\"/Ic:\\program files (x86)\\windows kits\\10\\include\\10.0.15063.0\\shared\" 
\"/Ic:\\program files (x86)\\windows kits\\10\\include\\10.0.15063.0\\um\" 
\"/Ic:\\program files (x86)\\windows kits\\10\\include\\10.0.15063.0\\winrt\" 
\"/Ic:\\program files (x86)\\windows kits\\10\\include\\10.0.15063.0\\cppwinrt\""
You must have the version 10.0.15063 Windows 10 SDK installed. This can be installed separately or by checking the appropriate box in the Visual Studio Installer.

该版本的SDK可以到这个页面上下载： https://developer.microsoft.com/en-us/windows/downloads/sdk-archive
