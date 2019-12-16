---
title: "Windows控制台上的UTF-8支持"
date: 2012-02-25T14:10:55+08:00
draft: false
tags: [golang]
---

最近在学习Go语言，发现调用一些库函数失败后的返回字符串，都是一些乱码，觉得奇怪。所以写了一个最简单的程序来复现这个问题：

```
func main(){ 
        _, err := os.Open("none-exist-file") 
        if err != nil{ 
                fmt.Println(err) 
        } 

} 
```
编译运行一切都正常，但是输出是乱码：
`pen none-exist-file: 绯荤粺鎵句笉鍒版寚瀹氱殑鏂囦欢銆? `

原因就是Go语言以UTF-8编码进行输出，而Windows控制台默认不支持UTF-8。解决方法就是更改code page，
使用命令 chcp，不加任何参数便显示当前的code page。把 ode page改为65001即可。另外还要把字体改为
Lucia Console。


这样的更改会导致其它一些不使用utf-8编码的程序出现乱码。


一些相关的link:
- https://groups.google.com/forum/#!topic/golang-nuts/WhpWNhBRMFM/discussion
- https://groups.google.com/group/golang-nuts/browse_thread/thread/cf9727737e5b4a00/d0e9d4dc975e875b?lnk=gst&q=messy+code#d0e9d4dc975e875b
- http://hi.baidu.com/edeed/blog/item/2e99a14440bd8884b2b7dcb1.html
- http://blog.codingnow.com/2008/12/utf-8_replacement.html