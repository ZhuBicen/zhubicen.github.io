---
title: "C++ Typename"
date: 2008-08-18T22:50:22+08:00
draft: false
tags: [C++]
---
```
#include<vector>
#include<numeric>
#include<iostream>
using namespace std;
template < class T>
class xlist
{
private:
    //typename vector<T>::iterator iter;
    vector<T>::iterator iter;
    vector<T> vi;
};

   
int main()
{
    xlist<int> x;
}
```
Under vc8:

1>c:\users\bzhu\documents\visual studio 2005\projects\t\t\t.cpp(9) : warning C4346: 'std::vector<T>::iterator' : dependent name is not a type
1>        prefix with 'typename' to indicate a type
1>        c:\users\bzhu\documents\visual studio 2005\projects\t\t\t.cpp(11) : see reference to class template instantiation 'xlist<T>' being compiled
1>c:\users\bzhu\documents\visual studio 2005\projects\t\t\t.cpp(9) : error C2146: syntax error : missing ';' before identifier 'iter'
1>c:\users\bzhu\documents\visual studio 2005\projects\t\t\t.cpp(9) : error C4430: missing type specifier - int assumed. Note: C++ does not support default-int

But under gcc it works well .Which is the c++ standard?