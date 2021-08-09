---
title: "Chrome扩展中的重要概念：Content Scripts"
date: 2009-08-28T17:47:39+08:00
draft: false
tags: [Chrome]
---

If you want your extension to interact with web pages, you need a content script.
一直在关注如何做Chrome的扩展，Chrome如何支持扩展。以前就看过Turorial中的HelloWorld，觉得这么简单，能干些什么啊。今天看了nkGestures的源码，然后又重新看了一下Chrome Extension的开发文档，觉得Content Script的概念比较重要，在此把我的理解记录下。

Chrome的扩展基本是基于Java Script的（当然可以通过其它方法调用C/C++代码），这样就为以后不同平台下的浏览器做好了准备，果然是Google。

Chrome允许在打开窗口，TAB或其它事件的情况下执行扩展中的代码，这部分代码就叫做：Content Scripts。Content Scripts are JavaScript files that run in the context of web pages. By using the standard Document Object Model (DOM), they can read details of the web pages the browser visits, or make changes to them.同时，Content Scripts 可以通过Channel和Extension进行通讯。

Content scripts can communicate with their parent extension using message passing. The content script opens a channel to the extension using the chrome.extension.connect() method and then sends messages back and forth to it. The messages can contain any valid JSON object (null, boolean, number, string, array, or object).

The parent extension can also open a channel to a content script in a given tab by callingchrome.tabs.connect(tabId).

When a channel is opened from a content script to an extension, the onConnect event is fired in all views in the extension. Any view can receive the event. The event contains a Port object, which can be used by the extension view to communicate back to the content scrip

