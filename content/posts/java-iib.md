---
title: "IIB(Instance Inilialization Block) in Java"
date: 2020-02-27T10:58:11+08:00
draft: false
tags: [Java]
---

```java
class Mugs {
    Mug mug1;
    Mug mug2;
    // 以下的代码块就是，IIB
    {
        mug1 = new Mug();
        mug2 = new Mug();
    }
    Mugs() {
        
    }
    Mugs(int i) {
        
    }
}
```



IIB 在构造函数执行之前执行，如果有多个构造函数，可以把多个构造函数的通用部分放入IIB中。只是有此看来，IIB并不是必须的，IIB 可以提取为通用的函数，在构造函数中调用也是可以的



但是，对于匿名内部类来说，对于成员的初始化则离不开 IIB 的支持，如下代码的第二层括号，就是IIB（第一层括号为创建匿名类）：



``` java
Map<String, String> doubleBraceMap  = new HashMap<String, String>() {{
    put("key1", "value1");
    put("key2", "value2");
}};
```

