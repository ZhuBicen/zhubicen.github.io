---
title: "C++17 Structure Binding"
date: 2019-09-23T20:19:23+08:00
draft: false
tags: [C++]
---

# 什么是Structured Binding

`Structured binding` 是 C++ 17 的一个新  feature, 他允许你绑定变量到实例或者构造列表中的元素。也可以这么说,  他允许开发人员同时声明多个变量, 并且有 `tuple` 或 `struct` 提供初始化值


他的主要目的是使代码更加易读。 像引用一样，`Structured Binding` 是一个已经存在对象的别名; 与引用不同的是 ` Structed Binding` 的类型可以不是引用



在 C++ 17 中可以写出如下的代码：

```C++
std::string func() {
    auto [succeeded, value] = get_some_pair();
    return value;
}
```

# Structed Binding 是如何工作的

编译器会把以上的代码翻译成如下的代码：

```C++
std::string func() {
    auto e = get_some_pair();
    auto& succeeded = e.first;
    auto& value = e.second;
    return value
}
```

# Structured Binding 与  `NVRO`
由于 `NRVO` 不能用在 `sub object` 因此，从性能方面考虑，如下的代码相比性能较好：

```C++
std::string func() {
    auto [succeeded, value] = get_some_pair();
    return std::move(value);
}

```


# 和 `std::vector` 一起使用

```C++
void f(std::array<int, 3> arr1, int (&arr2)[3], int (&arr3)[]) {
    auto [a1,b1,c1] = arr1;
    auto [a2,b2,c2] = arr2;
    auto [a3,b3,c3] = arr3;
}
```


[RVO, Copy elision](https://en.wikipedia.org/wiki/Copy_elision#Return_value_optimization)

https://www.youtube.com/watch?v=YePHP4aIc1g