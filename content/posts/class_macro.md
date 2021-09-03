---
title: "在编译期获取类名"
date: 2021-09-04T02:01:32+08:00
draft: false
tags: [C++]
---

最近参与的一个 C++ 项目里充斥着大量的 hardcode 类名和方法名的 log

```
log << "MyClass::MyMethod" << " calling ...";
```

总觉得应该有更好的办法（我的直觉常常是对的，这总能让我找到合适的办法做合适的事情）。印象中有 `__FUNCTION__`，但是如何获取类名呢，google 一番后发现了一个新的宏：`__PRETTY_FUNCTION__`，这个宏的确是 pretty，会打印出详细的函数名，包括所属名字空间，类名，方法名，还有类型签名。比如如果在 `main` 函数中使用，它会被处理成：`int main(int, char**)`。它甚至还可以打印出模版函数/类的类型参数，可谓是详细到不能再详细了（这或许是C++ 没有提供 `__CLASS__` 宏的原因吧）

```c++
// g++ -ggdb3 -O0 -std=c++11 -Wall -Wextra -pedantic -o main.out main.cpp
namespace N {
    class C {
        public:
            template <class T>
            static void f(int i) {
                (void)i;
                std::cout << "__func__            " << __func__ << std::endl
                          << "__FUNCTION__        " << __FUNCTION__ << std::endl
                          << "__PRETTY_FUNCTION__ " << __PRETTY_FUNCTION__ << std::endl;
            }
            template <class T>
            static void f(double f) {
                (void)f;
                std::cout << "__PRETTY_FUNCTION__ " << __PRETTY_FUNCTION__ << std::endl;
            }
    };
}

int main() {
    N::C::f<char>(1);
    N::C::f<void>(1.0);
}
```
打印出的内容为：
```
__func__            f
__FUNCTION__        f
__PRETTY_FUNCTION__ static void N::C::f(int) [with T = char]
__PRETTY_FUNCTION__ static void N::C::f(double) [with T = void]
```

但是对于我们的 log 来说似乎不需要这么详细的函数名，只需要类名+函数名就可以了。函数名可以用 `__FUNCTION__` 来实现。类名改如何实现呢，如果不是一定要宏的话，可以直接调用 `type(*this).name()` 就可以了，我觉得这个方法不错。但是仅仅是这样的方案还是不够完美，虽然这点运行时性能的损失微不足道（我觉得在大部分情况不应该成为问题），但是为了避免别的同事有这样的顾虑，能不能基于`__PRETTY_FUNCTION__` 来实现一个的宏呢

没有做不到只有你想不到，还真有，上代码（*今天的主角登场了*）：


``` C++
constexpr std::string_view method_name(const char* s)
{
    std::string_view prettyFunction(s);
    size_t bracket = prettyFunction.rfind("(");
    size_t space = prettyFunction.rfind(" ", bracket) + 1;
    return prettyFunction.substr(space, bracket - space);
}
#define __METHOD_NAME__ method_name(__PRETTY_FUNCTION__)
```

咋一看没什么稀奇，不就是把`__PRETTY_FUNCTION__` 处理了一下嘛。非也，有几个重要的点
+ `constexpr` 这意味着函数返回一个常量的表达式，因此表达式的值需要在编译期确定。在编译期 `methodName` 方法已经被调用了，没有运行时的开销
+ `std::string_view` 这个类的大部分方法都可以在编译期运行（大多数方法是`constexpr`属性)
+ 把 `std::string_view` 换成 `std::string` 是不行的，会有编译错误。不是任何代码都可以放在编译期执行的

>> 需要注意的是 `constexpr` 是在 C++11 引入的，`std::string_view` 是在 C++17 引入的。C++ 版本太低的话，能升就升吧

如此这般再打 log 的话就爽了，至少你 refactor 函数名的时候，不用担心不一致了

```
log << __METHOD_NAME__ << " calling ...";
```

上面的代码基本来自于[这里](https://stackoverflow.com/questions/1666802/is-there-a-class-macro-in-c)


