---
title: "C++20 Allow lambda capture [=, this]"
date: 2020-03-03T17:12:17+08:00
draft: true
tags: [C++20]
---


在 C++20 之前 `[=, this]` 是视为语法错误的，有些编译器视为 Warning。
C++20 允许了这种形式的出现。其等价于 `[=]`，只是显示的指定了 `this`。

`
以 copy 的形式不会 this 指针， 因为 `this` 是指针，相当于 以 引用的形式捕获了 `*this`


http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2017/p0409r2.html

![1583228112256](c++20_allow_lamdba_capture_this.assets/1583228112256.png)



 https://docs.microsoft.com/en-us/cpp/cpp/lambda-expressions-in-cpp?view=vs-2019 



```C++
struct S { void f(int i); };

void S::f(int i) {
    [&, i]{};      // OK
    [&, &i]{};     // ERROR: i preceded by & when & is the default
    [=, this]{};   // ERROR: this when = is the default
    [=, *this]{ }; // OK: captures this by value. See below.
    [i, i]{};      // ERROR: i repeated
}
```

