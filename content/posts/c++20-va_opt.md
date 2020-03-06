---
title: "C++20 学习笔记之 VA_OPT"
date: 2020-03-06T10:20:50+08:00
draft: false
tags: [C++20]
---

`VA_OPT` 的特性与宏的可选参数(`VA_ARGS`)有关，经常配合起来使用，比如宏：

```c
#define LOG(msg, ...) printf(msg, __VA_ARGS__)
LOG("hello %d\n", 0);   // => printf("hello %d", n)
```

如上对 `LOG` 宏的调用是没有问题的，但是如果我们想更灵活一点，如下的调用，就会出现问题：

```c
LOG("hello world");  // => printf("hello world,")
```
可以看倒展开后多了一个**逗号**，这样将会导致编译错误。

现在就是`VA_OPT` 出场的时候了，`LOG` 宏可以做如下改造，以使**逗号** 变为可选

```c
#define LOG(msg, ...) printf(msg __VA_OPT__(,) __VA_ARGS__)
```

`__VA_OPT__` 回根据 `__VA_ARGS_` 是否为空来决定 **逗号**  是否出现，因此可以解决`LOG("hello world")` 的调用问题。完整代码如下：

```c
#include <stdio.h>
#define LOG(msg, ...) printf(msg __VA_OPT__(,) __VA_ARGS__)
// #define LOG(msg, ...) printf(msg, __VA_ARGS__)

int main() {
    LOG("hello world\n");   // => printf("hello world")
    LOG("hello world\n", );  // => printf("hello world")
    LOG("hello %d\n", 0);   // => printf("hello %d", n)

    // printf("Hello world",);

    return 0;
}
```

代码已经放在 github 上： https://github.com/ZhuBicen/Clang-in-Docker/blob/master/Tests/va_opt.cpp



另外，`VA_OPT` 的另外一个例子：

```C++
#define SDEF(sname, S, ...) S sname __VA_OPT__(= { __VA_ARGS__ })

SDEF(foo);           // => S foo;
SDEF(bar, 1, 2, 3);  // => S foo = { 1, 2, 3 };
```



最后，对于 C++ 的大多数情况，应该避免对于宏的使用