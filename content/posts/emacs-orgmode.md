---
title: "在 Emacs 中使用 org-mode"
date: 2012-03-09T17:52:27+08:00
draft: false
tags: [emacs]
---

使用emacs24默认的包管理机制使得安装一些扩展包异常的方便。：D
```lisp
  (package-initialize)
  (add-to-list 'package-archives '("marmalade" . "http://marmalade-repo.org/packages/"))
  ;;如果你使用代理的话，请uncommnet
  ;;(setq url-proxy-services '(("http" . "127.0.0.1:8080")))
  
```

然后可以使用package-list命令，此时emacs会生成一个buffer，在其中列出所
有可用的软件包，找到org2blog，用鼠标点进去，再点击Install即可。：D，所
有这一切都是无痛的.

最后你可以到org2blog的页面上查看如何使用。[[https://github.com/punchagan/org2blog][org2blog]]

希望自己以后要多总结！：D 

加油 :-)