---
title: "ES6 中的 const"
date: 2019-12-17T22:32:48+08:00
draft: false
tags: [JavaScript]
---

Java Script 中的 const 变量，在执行的过程中不能重新赋值，但是可以对原有值更改。

比如：

```JavaScript
const constVar = {};
constVar.newProperty = "newValue";
```

这个初看起来挺奇怪的，但是 `C++` 其实也是如此。由于 `JavaScript` 中的变量是引用，所以那 `C++` 中的引用进行对比的话，两者的语义是相近的，但是不完全一致

如下代码是可以的：

```c++
MyClass myObject;
const MyClass& constObject = myObject;
myObject.a = 1;
```

但是在`C++` 中不可直接使用 `const reference` 来改变值的属性：

```C++
MyClass myObject;
const MyClass& constObject = myObject;

constObject.a = 2;// 这样是不行的: expression must be a modifiable value
```

