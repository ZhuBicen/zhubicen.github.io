---
title: "Override VS Overload"
date: 2021-08-02T20:35:49+08:00
draft: false
tags: [C++]
---

同事问了我一个问题，google 了一番后，答案指向了 Stack Overflow:
https://stackoverflow.com/questions/35870081/c-base-class-function-overloading-in-a-subclass

又在老爷子的书上找了下源头：

## Overload 覆盖（子类覆盖父类)

覆盖的条件：

1. 必须是虚函数
2. 函数名和函数类型（参数类型）必须一致


## Override 

Override 不能跨作用域


```C++
struct Base {
    void f(int);
};
struct Derived : Base {
    void f(double);
};
void use(Derived d)
{
    d.f(1); // call Derived::f(double)
    Base& br = d
    br.f(1); // call Base::f(int)
}
```

但是可以使用 “Using” 来改变方法的作用域，比如：


``` C++
struct D2 : Base {
    using Base::f; // bring all fs from Base into D2
    void f(double);
};

void use2(D2 d)
{
    d.f(1); // call D2::f(int), that is, Base::f(int)
    Base& br = d
    br.f(1); // call Base::f(int)
}
```

参考：
1. The C++ programming langurage forth edition.pdf $20.3.5