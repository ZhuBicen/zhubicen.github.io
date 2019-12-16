---
title: "使用GYP生成MSVS工程文件~0"
date: 2014-03-25T01:22:18+08:00
draft: false
tags: [C++]
---

GYP是Google出品的新一代C++构建系统，被广泛的用于Chrome的构建中。

准备工作：

1. svn checkout http://gyp.googlecode.com/svn/trunk/ gyp-read-only
2. 设置gyp-read-only目录到PATH
3. set GYP_GENERATORS=msvs

由于GYP在windows下默认使用ninja来编译，所以要生成MSVS工程文件，需要设置GYP_GENERATOR为msvs。如果你安装了多个版本的Visual Studio，也可以指定MSVS的版本，比如 set GYP_MSVS_VERSION=2013，目前GYP支持的版本包括2008到2013的所有版本，包括express版，如果是express版在版本号后加e即可。比如2013e为Visual Studio 2013 express版。

创建如下的GYP文件，program.gyp，
```
{
  'targets': [
    {
      'target_name': 'program',
      'type': 'executable',
      'sources': [
        'program.cpp',
      ],
    },
  ],
}
```
运行如下命令，并确保当前目录下存在 program.cpp
```
gyp --depth=. program.gyp
```
检查当前目录，会生成如下的文件，
`program.sln`和`program.vcxproj`，使用 Visual Studio 打开 `program.sln` 即可。

一般来说，使用MSVS创建的项目都会有Debug和Release两种模式，但是目前生成的项目文件只有一种Default配置，无法在里面设置断点。 下一篇文章解决此问题。