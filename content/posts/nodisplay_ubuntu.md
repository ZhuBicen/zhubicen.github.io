---
title: "远程使用没有显示器的 Ubuntu"
date: 2021-11-19T21:41:46+08:00
draft: false
tags: []
---

朋友送了一台 Mac Mini （2012 年底版），以前有一个 Windows 台式机，还有两个显示器，就这么搭配着用，不爽的地方是总是要在两个系统之间要切换键盘鼠标，还有鼠标操作总是感觉有些卡顿。所以用的机会不多。今年的买了一个 iMac，由于想精简做面，就把显示器还有台式机都处理掉了。随之发现一个问题，由于 iMAC 是M1 芯片的，所以在上面使用 Docker 时也不方便，因为 arm 版的 Docker Image 也是不多

买了 iMac 后，Mac Mini 使用的机会就更少了，想着把它装个 Linux，然后远程使用，由于还想使用 gstreamer，最好还是可以使用远程桌面，不能只是 SSH。由于家中没有显示器，只能把 Mac Mini 带到公司装好系统，然后拿回家使用。在此过程中经历和解决了一些问题，特此记录

1. Ubuntu 自带远程桌面，在设置中可以找到，称为 Screen Sharing，因此不需要安装配置另外的 VNC server。在设置中 enable 后需要打开防火墙端口 5900。
2. 一般的智能终端（包括这里的 Mac Mini Ubuntu）都会自动连接已经记住密码的 Wifi，利用这个特性就可以在公司装好系统连接手机热点模拟的 Wi-Fi（使用和家中同样的 Wi-Fi 名称和密码），然后带回家中后就可以自动连接家中的 Wi-Fi。
3. Ubuntu 的 Screen Sharing 还需要勾选网络，有线网络或者对应的 Wifi，因此在网络变动后（比如从公司拿回家中），只有 IP 的话还是连接不上的。需要 Enable 对应的网络（Wi-Fi)
5. 另外最奇葩的一点，如果 Mac Mini 没有连接外接显示器，使用 Screen Sharing 连接后看到的是一个黑屏。要解决这个问题需要安装一个虚拟的显示器（fake display)。经测试自己安装的 VNC server（使用 Xfce 桌面套件等，没有此问题）
6. 通过家中的路由器管理 Web 找到 Mac Mini 的IP

解决了以上问题后就可以完美的把 Mac Mini 带回家中使用，并且可以远程到自带的 Ubuntu 桌面

![image-20211111215803535](/Users/bzhu/Project/quickstart/content/posts/nodisplay_ubuntu.assets/image-20211111215803535.png)

有用的链接：
[添加虚拟显示器](https://askubuntu.com/questions/453109/add-fake-display-when-no-monitor-is-plugged-in)
[VNC viewer 连接不上的问题](https://askubuntu.com/questions/1126714/vnc-viewer-unable-to-connect-encryption-issue?newreg=d90a4634e7224b12a64ff8d9c4c59b14)
