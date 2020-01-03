---
title: "Chrome 未请求 Favicon"
date: 2020-01-03T20:51:22+08:00
draft: false
tags: [Chrome]
---

发现 Chrome 并不是每次都请求 favicon，硬性重新加载即可

![image-20200103205715291](/Users/bzhu/Library/Application Support/typora-user-images/image-20200103205715291.png)

要为整个 Website 设定 Favicon 需要在文件服务器的根目录放置文件 favicon.ico， 以使 `http://you.host.name/favicon.ico` 可用。如果要为单个网页设定 Favicon， 可以在对应的 html 文件 head 中加入：

```

<head>
   <link rel="shortcut icon" href="http://example.com/favicon.ico" type="image/vnd.microsoft.icon" />
</head>
```
