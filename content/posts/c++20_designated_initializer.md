---
title: "C++20 学习笔记之 designated initializer"
date: 2020-03-28T20:18:30+08:00
draft: false
tags: [C++20]
---



```c++
struct A {
	string a;
	int b = 42;
	int c = -1;
};

```

`A {.c=21} ` 会执行如下的步骤：

1. 初始化字符串 `a` 为空
2. 初始化`b` 为 42
3. 初始化`c` 为 21

指定的标识符需要是非静态的数据成员，并且需要和声明的顺序一致

```C++
struct A 
{ 
  int x;
  int y;
  int z;
};

A a{.y = 2, .x = 1}; // 这样使不行的，x 必须出现在 y 的前面
A b{.x = 1, .z = 2}; // OK y 将被初始化为 0
```



另外指定的初始化器还可以对 union 进行初始化，但是只允许指定其中一个标识符：

```c++
union u 
{ 
  int a; 
  const char* b; 
};

u f = { .b = "asdf" }; // 可行
u g = { .a = 1, .b = "asdf" }; // 错误，因为使用了多个指定标识符
```



C++ 20 中的 `designated initializer ` 和 C 语言中的是不兼容的，比如 C++ 的对于变量的顺序有要求，不支持数组元素，不支持嵌套，不支持混合模式

```C++
struct A 
{ 
  int x, y; 
};
struct B 
{ 
  struct A a; 
};

struct A a = {.y = 1, .x = 2}; // 顺序不对
int arr[3] = {[1] = 5};        // C++ 不支持数组元素
struct B b = {.a.x = 0};       // C++ 不支持嵌套元素
struct A a = {.x = 1, 2};      // C++ 不支持混合模式

```

而以上的几种使用形式在C语言中都是合法的

