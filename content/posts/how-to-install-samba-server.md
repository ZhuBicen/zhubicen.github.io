---
title: "How to Install Samba Server"
date: 2007-01-25T22:36:03+08:00
draft: true
tags: [Linux]
---

apt-get install samba

uncomment the command in /etc/samba/smb.conf 

Security = share

add a file to share ,append a share node in smb.conf .and the folder should be 777

在Ubuntu 6.10 Desktop和Windows XP之间设置文件共享 

在Ubuntu下打开“系统-系统管理-共享的文件夹”，进行相应的设置。在WinXP下打开Ubuntu的IP，可以看到文件夹，但是并不能够访问。 
用如下命令修改smb.conf： 

sudo gedit /etc/samba/smb.conf 

将security=user那一行前的注释符";"去掉，然后把user改为share，这样可以实现匿名访问。 
再用如下命令重启samba： 

sudo /etc/init.d/samba restart 

在WinXP，可以访问文件夹了，但是只读，仍然无法写入。即使在Ubuntu的GUI下去掉了只读属性也是如此。 
最后修改文件夹的权限： 

chmod 777 (文件夹名称) 