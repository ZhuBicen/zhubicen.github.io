---
title: "解读 C++ 指针声明"
date: 2009-06-27T10:33:33+08:00
draft: false
tags: [C++]
---
```
char *const cp ; // const pointer to char 
char const * pc ; // pointer to const char
const char * pc2 ; // pointer to const char 
Some people find it helpful to read such declarations right to left. For example, "cp is a const pointer to a char" and "pc2 is a pointer to a char const ."
```
如何复杂的表达式声明，掌握了方法一样可以读出来。首先，()和[]具有比*更高的优先级。因此先读[]，读作 array of，*读作pointer to。当然()具有最高的优先级，如果（）里没有内容类似 returntype (* function)()的，读作function。

`char (*(*x())[])()`，从里面往外读 , x is a function returning pointer to D，读到这里可以进行一个简化， `char(*D)[]()`，继续读，D is array of function returning char。因此： function returning pointer to array[] of  pointer to function returning char

```
char (*(*x[3])())[5]  --->> char(*D)())[5]，x is array3 of pointer to , D is function return char[5]。
```

这是我比较好理解的方法，写的也比较粗糙。更详细请参考这里：

http://blog.chinaunix.net/u/12783/showart_378340.html
```
int
int *
int *[3]
int (*)[]
int *()
int (*[])(void)
```
