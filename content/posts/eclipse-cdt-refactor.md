---
title: "使用 Eclipse Cdt 重构 C++ 代码"
date: 2008-11-12T23:04:07+08:00
draft: false
tags: [C++]
---
一年前用ＣＤＴ时还发现其经常会失反应，时至今日又试用了一下，发现已有很大改善。并且在其官网上可以下载到，专门为Ｃ＋＋程序员而打包的Eclipse，这样我们就不用再自己去配置插件了：）。虽然Eclipse有些笨重，但是相对安装一下就要近半个小时的ＶＳ而言，已然轻便多了。

ＣＤＴ的重构主要由以下几个功能，相对对ＪＡＶＡ的支持，差远了。这可能是由于Ｃ＋＋本身的复杂性引起的。

```
      Rename
      Extract Constant
      Extract Function
      Hide Method（这个我还末试用）
      Implement Method（这个是可以实现类的头文件里声明的函数，好像与继承无关。。）
      Generate Getter and Setter（这个功能对Ｃ＋＋啥用吧：），并且实现的也有点傻）
```

![image-20191216230612109](eclipse-cdt-refactor.assets/image-20191216230612109.png)
![image-20191216230639462](eclipse-cdt-refactor.assets/image-20191216230639462.png)
