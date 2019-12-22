---
title: "Java Functional Interface"
date: 2019-12-22T01:06:03+08:00
draft: false

---

请看如下的 Java 代码：
```Java
class MyClass {
public setCalback(Runnable doneCallback) {
		if (something is done) {
			doneCallback();
		}
	}
};
MyClass mc = new MyClass();
mc.setCallback(() -> {
	// do other things
})
```
> 为什么 `lambda` 可以实现`Runnable` 接口？

`Runnable` 的代码如下， 可以看到他是一个函数式接口 `FunctionalInterface`，类似的函数是接口，在 Java中还有很多，比如 `ActionLisener` 和 `Comparator` 

```Java
package java.lang;

@FunctionalInterface
public interface Runnable {
    void run();
}
```

## 什么是函数是接口？
函数式接口是一种只有一种方法的接口。具有相同签名的 `lambda` 函数可以转化为这种接口。

> A functional interface is an interface that contains only one abstract method. They can have only one functionality to exhibit. From Java 8 onwards, lambda expressions can be used to represent the instance of a functional interface. A functional interface can have any number of default methods. Runnable, ActionListener, Comparable are some of the examples of functional interfaces.


