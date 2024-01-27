---
title: "Xpath Syntax"
date: 2024-01-27T19:34:06+08:00
draft: false
tags: [xpath]
---

XPath -> XML Path Language
你还别小瞧，XPath 是一门语言，有 200 个内建的函数，这些函数可以操作字符串，数字，布尔，日期，时间，结点等。 它有三个版本，目前是 3.0 版本

在 XPath 语言中有一下几种结点：
element, attribute, text, namespace, processing-instruction(这是什么？)，注释和 root element

```xml
<?xml version="1.0" encoding="UTF-8"?>

<bookstore>
  <book>
    <title lang="en">Harry Potter</title>
    <author>J K. Rowling</author>
    <year>2005</year>
    <price>29.99</price>
  </book>
</bookstore>
```

`bookstore` 为 root element node，`book` 为 element node
`lang="en"`为 (attribute node)，`book` 为 `title, author, year, price` 的父节点。

例子

```
# 选择  bookstore 下的第一个 book 结点
/bookstore/book[1]
# 最后一个
/bookstore/book[last()]
# 倒数第二个
/bookstore/book[last()-1]
# 选取前两个
/bookstore/book[position()<3]
# 选择价格大于 35
/bookstore/book[price>35.00]
# 选择价格大于 35 的 title
/bookstore/book[price>35.00]/title
```

[Xpath cheatsheet](https://devhints.io/xpath)

可以参考以上的语法，把需要提取的 `xml` 和 自己写的 `xpath` 放到 [在线 1](http://xpather.com/) 或者 [在线 2](http://xpather.com/) 测试，提前检查一些语法的错误
