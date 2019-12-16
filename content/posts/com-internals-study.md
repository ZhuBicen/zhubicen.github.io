---
title: "Com Internals Study"
date: 2009-09-01T12:58:46+08:00
draft: false
tags: [C++]
---

**COM技术内幕读书笔记** 

第１章ＣＯＭ历史

**第２章接口（用Ｃ＋＋的纯抽象基类来实现ＣＯＭ接口）**

**Ｑ１：什么是接口？**

DLL的接口就是它输出的那些函数，Ｃ＋＋类的接口就是则是该类的一个成员函数集。ＣＯＭ接口也涉及一组由组件实现并提供给客户使用的函数，是一个包含一个函数指针数组的内存结构。在Ｃ＋＋中可以用抽象基类（不是抽象函数）来实现接口，一个组件可以实现任意多的接口，Ｃ＋＋可以用抽象基类的多重继承来实现。

以＝０标记的虚拟函数称为纯虚函数，仅包含纯虚函数的类称为纯抽象基类。

标准调用约定与Ｃ调用约定（__stdcall 和__cdecl)固定参数的Win32API调用采用的都是__stdcall，可变参数调用采用__cdecl

COM并没有要求一个Ｃ＋＋类同一个ＣＯＭ相对应。实际上一个组件可以由多个Ｃ＋＋类来实现。

ＣＯＭ并没有要求实现某个接口的类必须从那个接口继承，这是因为客户并不需要了解ＣＯＭ组件的继承关系。对接口的继承只不是一种实现细节而已。

**Ｑ２虚拟函数表**

struct IX{

   virtual void __stdcall Fx1() = 0;

   virtual void __stdcall Fx2() = 0;

   virtual void __stdcall Fx3() = 0;

   virtual void __stdcall Fx4() = 0;

}

幸运的是，所有与Windows兼容的Ｃ＋＋编译器都能生成ＣＯＭ可以使用的正确的vtbl。

Ｃ＋＋编译器为纯抽象基类所生成的内存结构同ＣＯＭ接口所要求的内存结构是相同的。



ＣＯＭ需要所有的接口都支持三个函数，这三个函数必须是接口的vtbl中的前三个函数(用来做接口查询）。



**第３章，QueryInterface函数**

interface IUnkown

{

   virtual HRESULT __stdcall QueryInterface(const IID& idd, void** ppv) = 0;

   virtual ULONG __stdcall AddRef() = 0;

   virtual ULONG __stdcall Release() = 0;

};

IUnkown* CreateInstance();

其中ＩＩＤ为接口标识符

QueryInterface是COM最为重要的部分，这主要是因为一个组件实际上是由QueryInterface定义的，



**第４章，引用计数**



该章介绍了，如何使用AddRef与Release的一些规则，如果借用了智能指针，这些是很容易解决地

**第５章　动态连接库** 

该章使用用ＤＬＬ及ＣＯＭ标准建立了一种可以在运行时将组件和客户链接起来的架构。

ＤＬＬ：导出一个CreateInstance的函数，该函数返回一个IUnkown指针，这是ＤＬＬ接口的全部。

客户：　LoadLibary后调用CreateInstance获取一个IUnknow指针，然后根据该接口，查询其它的接口。

此外，Iface.h中定义了ＣＯＭ组件（使用ＤＬＬ实现）所实现的接口信息。The content of Iface.h is as follows:

```

*interface IX: IUnknown*

*{*

   *virtual void __stdcall Fx() = 0;*

*};*

*interface IY: IUnknown*

*{*

   *virtual void __stdcall Fy() = 0;*

*};*

*interface IZ: IUnknown*

*{*

   *virtual void __stdcall Fz() = 0;*

*};*

*extern "C"*

*{*

   *extern const IID IID_IX;*

   *extern const IID IID_IY;*

   *extern const IID IID_IZ;*

*}*


```


**第6章 关于HRESULT, GUID，注册表和其它细节**



通过HRESULT值返回正确或者错误的代码，，，不太清楚，，，

GUID_DEFIN宏不太清楚。

COM库的初始化，CoInitialize, CoUninitialize。

COM库中的void* CoTaskMemAlloc(ULONG cb)和 void CoTaskMemFree(void* pv);可以在不同的进程之间，多线程之间分配与释放内存，，强悍！！

COM库定义了一些宏，使用C/C++以同样的文式定义COM接口

第7章 类厂

客户可以先择以不同的方式创建组件，在**进程中，在本地，在远程服务器**。。。

类厂也是COM库的一部分。现在感觉微软COM包括两个部分，COM标准及微软COM库（是为方便操作COM而生成）

第8章 组件复用 包容与聚合

继承包括实现继承与接口继承。实现继承指的是类继承其基类的代码或实现。接口继承是类继承与基类的类型或接口。COM不支持实现继承，我们可以用组件包容来完全模拟实现继承。

包容和聚合实际上是使一个组件使用另外一个组件的一种技术。对于这两个组件，作者分别将其称为是外部和内部组件。在包容的情况下，外部组件包容内部组件；而在聚合的情况下，刚称外部组件聚合内部组件。

聚合是包容上特例。当一个外部组件聚合了某个内部组件的一个接口时，它并没有像包容一样实现此接口并没有明确地将请示转发给内部组件。相反，外部组件将直接把内部组件的接口指针返回给客户。

**聚合是的目的就是使客户确信内部组件所使用的接口是由外部组件实现的。被聚合的组件需要知道外部接口的IUnknow指针，并知道自己是被聚合的。**