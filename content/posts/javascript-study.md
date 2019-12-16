---
title: "Javascript 学习笔记"
date: 2009-06-05T23:29:38+08:00
draft: false
tags: [JavaScript]
---

from http://www.howtocreate.co.uk/tutorials/javascript/

## 定义变量

JS里变量的类型并不是非常重要，变量的类型可以在运行期改变。基础类型有，character，string，integer，float（or double)，boolean, function, object, array, undefined, null. Interger and float are grouped as ‘number’. Charater and string are grouped as string. array属于对象类型。

变量的定义：

```
var myFish = ‘A fish swam in the river’ 
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

对象定义：

```
var o1 = new Object(); 
var o2={myFirstProperty:1, myNextProperty:’hi’,etc};
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

数组的使用：

```
var nameOfArray = new Array();
var nameOfArray = new Array('content_of_first_cell','and_the_second',8,'blah','etc');
var nameOfArray = ['content_of_first_cell','and_the_second',8,'blah','etc'];
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

```
同时JavaScript可以定义多维数组：
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

```
var nameOfArray = new Array(new Array(1,2,3,4),'hello',['a','b','c']);
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

其它类型：Date,Image,Option,RegExp

If at any time you refer to a variable that does not exist, you will cause an error. However, JavaScript allows you to refer to one level of undefined child properties.

JavaScript中也有值类型和引用类型的概念，所有的基本类型为值类型，所有的object类型为引用类型。

```
var myNewVariable = myOldVariable;
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

If myOldVariable was already defined as a string, number, boolean, null or undefined, it would simply have copied myOldVariable and 'pasted' it into myNewVariable. If the new variable were changed (for example, using `myNewVariable = 'some new value';`)myOldVariable retains its old value.

If, on the other hand, myOldVariable was already defined as a function, array, object or option, myNewVariable would have been created as a pointer to myOldVariable. The children of myOldVariable would now also be the children of myNewVariable. If the new variable were changed (for example, using `myNewVariable = 'some new value';`), it would only alter the value of myNewVariable so that it no longer references the same contents as myOldVariable. However, changing the child properties of myNewVariable will change the properties of the object it references, and as a result, myOldVariable would also see the change:

```
var myOldVariable = new Object();
var myNewVariable = myOldVariable;
myNewVariable.newChild = 'Hello';
alert(myOldVariable.newChild);
//This will alert 'Hello'
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

## 使用函数

The following is an example of using the return statement to return a variable from a function:

```
function appendComment(passvar) {
  //in this case, I have passed a string variable and I return
  //a string variable. I could return any variable type I want.
  passvar += ' without your help';
  return passvar;
}

var myString = appendComment('I did it');
//myString is now 'I did it without your help'
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)



Functions are called using one of these:

```
nameOfFunction(listOfVariables); 
window.nameOfFunction(listOfVariables); 
object.onEventName = nameOfFunction;
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

## DOM操作

http://www.howtocreate.co.uk/tutorials/javascript/dombasics