---
title: "Java Charset"
date: 2019-12-23T11:32:30+08:00
draft: false
tags: [Java]
---

在给公司做抽奖程序的时候，做 了这么一个功能，从 csv 读取内容放到数据库，在 IDE 里测试的好好的。拿到命令行运行后，插入数据库就是乱码。

后来发现是读取文本的内容的时候没有设置UTF8，但是为什么 IDE 是好的？IDE（包括 Eclipse 和 IntelliJ）里下直接就是 UTF8，所以不用设置Java IDE 包括 Eclipse 和 IDEA 默认的字符集都是 UTF-8，而系统本身的字符集（在使用命令行启动应用程序的时候），对于中文操作系统来说一般是 `GBK`，比如下面的代码，在 IDE 中运行的时候会打印 `UTF-8`，在 Windows 命令行运行的时候打印 `GBK`

``` Java
public class Application {
    public static void main(String[] args) throws InterruptedException {
        System.out.println(Charset.defaultCharset());
        System.out.println(new OutputStreamWriter(new ByteArrayOutputStream()).getEncoding());
        System.out.println(System.getProperty("file.encoding"));
    }
}
```

