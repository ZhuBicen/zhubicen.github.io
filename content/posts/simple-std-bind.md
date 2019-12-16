---
title: "一个简化版的 std::bind 实现"
date: 2012-12-02T17:22:04+08:00
draft: false
tags: [C++]
---
```C++
#include <iostream>
#include <algorithm>
#include <vector>

namespace {
  class placeholder1 { };
  class placeholder2 { };
  placeholder1 _1;
  placeholder2 _2;
}

template <typename R,typename T, typename Arg1, typename Arg2>
class simple_bind_t {
  typedef R (T::*fn)(Arg1, Arg2);
  fn fn_;
  T t_;
public:

  simple_bind_t(fn f,const T& t):fn_(f),t_(t) {}

  R operator()(Arg1 a, Arg2 b) {
    return (t_.*fn_)(a, b);
  }
};


template <typename R,typename T, typename Arg1, typename Arg2>
class simple_bind_t2 {
  typedef R (T::*fn)(Arg1, Arg2);
  fn fn_;
  T t_;
public:

  simple_bind_t2(fn f,const T& t):fn_(f),t_(t) {}

  R operator()(Arg2 a, Arg1 b) {
    return (t_.*fn_)(b, a);
  }
};


template <typename R, typename T, typename Arg1, typename Arg2>
simple_bind_t <R,T,Arg1, Arg2> 
simple_bind(R (T::*fn)(Arg1, Arg2), const T& t, const placeholder1&, const placeholder2&) {
    return simple_bind_t<R,T,Arg1, Arg2>(fn,t);
}

template <typename R, typename T, typename Arg1, typename Arg2>
simple_bind_t2 <R,T,Arg1, Arg2> 
simple_bind(R (T::*fn)(Arg1, Arg2), const T& t, const placeholder2&, const placeholder1&) {
    return simple_bind_t2<R,T,Arg1, Arg2>(fn,t);
}

class Test {
public:
  void do_stuff(const std::vector<int>& v, int i) {
    std::copy(v.begin(),v.end(), std::ostream_iterator<int>(std::cout," "));
    std::cout << " i = " << i << std::endl;
  }
};

int main() {
  Test t;
  std::vector<int> vec;
  vec.push_back(42);
  vec.push_back(32);
  simple_bind(&Test::do_stuff, t, _1, _2)(vec, 999);
  simple_bind(&Test::do_stuff, t, _2, _1)(999, vec);
}
```