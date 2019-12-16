---
title: "收不到TaskbarButtonCreated消息"
date: 2012-03-03T16:06:37+08:00
draft: false
tags: [C++]
---

一个WTL对话框小程序，想为其加上Win7特有的任务栏上的进度条功能，死活不行，注册消息
`RegisterWindowMessage ( _T("TaskbarButtonCreated") )`;也是成功的，但是无论如何都收不到消息。

重新用Wizard建立了一个最简单的小程序，发现是可以的。对比了半天，慢慢缩小问题的范围，折磨了几个小时，最后发现原来我的对话框有一个属性是，`WS_EX_TOOLWINDOW`, 所以把此属性去掉即可。


```
diff -r fd2da7b96050 Ling/Ling.rc
--- a/Ling/Ling.rc      Sat Mar 03 13:57:18 2012 +0800
+++ b/Ling/Ling.rc      Sat Mar 03 16:00:40 2012 +0800
@@ -89,7 +89,6 @@
 LANGUAGE LANG_NEUTRAL, SUBLANG_NEUTRAL
 IDD_MAINDLG DIALOG 0, 0, 351, 28
 STYLE DS_3DLOOK | DS_CENTER | DS_SHELLFONT | WS_BORDER | WS_VISIBLE | WS_POPUP | WS_SYSMENU
-EXSTYLE WS_EX_TOOLWINDOW
 FONT 8, "Ms Shell Dlg"
 {
     CONTROL         "退出", IDOK, WC_BUTTON, WS_TABSTOP | WS_TABSTOP | BS_OWNERDRAW, 335, 0, 14, 14
diff -r fd2da7b96050 Ling/Release/Ling.exe
Binary file Ling/Release/Ling.exe has changed
```
