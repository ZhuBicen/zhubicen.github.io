---
title: "一个 Java 项目的总结"
date: 2009-08-26T17:42:31+08:00
draft: false
tags: [C++]
---

又一次要搞一下 Java，每次都要装下 JDK，设置下环境变量，编译一个 HelloWorld来运行一下。下面再记录一次：

1. 设置JAVA_HOME为JDK的根目录

2. 把%JAVA_HOME%bin添加进path

3. 设置classpath为 “.;%JAVA_HOME%/lib/dt.jar;%JAVA_HOME%/lib/tools.jar”

至此编译HelloWorld应该不成问题了。

由于这是一个关于Web`struts2+Hibernate+Spring`的项目

首先我通读了 CoreJsp一书。最简单的JSP程序：




A simple date example

```
The time on the server is
<%= new java.util.Date() %> 
```



之后该书用了大量的篇幅讲解了JSP的标签用法及与ServerLet相关的一些概念。作为标准，JAVA Web程序的根目录下要有一个WEB-INF文件夹，The WEB-INF directory contains the web.xml file as well as two different directories: classes and lib. classes contains Servlet and utility classes needed by the application. lib contains JAR files that might contain Servlets, JavaBeans, and utility classes that might be useful to the application. The web.xml file, commonly known as the deployment descriptor, contains configuration and deployment information about the Web application. It can contain information about the application in general, such as MIME type mappings, error pages, session configuration, and security. It can contain information about definitions and mappings to both JSP pages and Servlets. Another type of information that can be contained within the deployment descriptor is Servlet context initialization parameters.



然后就是研究一下struts2

Your browser sends to the web server a request for the URL http://localhost:8080/tutorial/HelloWorld.action.

1. The container receives from the web server a request for the resource HelloWorld.action. According to the settings loaded from the [web.xml](http://struts.apache.org/webxml.html), the container finds that all requests are being routed toorg.apache.struts2.dispatcher.FilterDispatcher, including the *.action requests. The FilterDispatcher is the entry point into the framework.
2. The framework looks for an action mapping named "HelloWorld", and it finds that this mapping corresponds to the class "HelloWorld". The framework instantiates the Action and calls the Action's executemethod.
3. The execute method sets the message and returns SUCCESS. The framework checks the action mapping to see what page to load if SUCCESS is returned. The framework tells the container to render as the response to the request, the resource HelloWorld.jsp.
4. As the page HelloWorld.jsp is being processed, the  tag calls the getter getMessage of the HelloWorld Action, and the tag merges into the response the value of the message.
5. A pure HMTL response is sent back to the browser.

For detailed information on Struts 2 architecture see [Big Picture](http://struts.apache.org/big-picture.html).

最后又看了Hibernate，和Spring的tutorial，对这些技术有了一些了解。