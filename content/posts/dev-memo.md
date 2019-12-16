---
title: "Windows X-Server 使用记录
date: 2010-12-23T20:02:24+08:00
draft: false
tags: [Linux]
---

ssh-host-config,ssh-user-config用来配置cygwin的ssh server.
完了之后中启动总是失败，查无果，无奈，，看到windows service的log后，发现问题了，原来是我的path中放入了一个cygwin1.dll。。这真是个个别情况。到网上查成了一种习惯，有的问题还是要自己查找原因。看LOG很重要。。
Cygrunsvr.exe可以操作Windows服务。

如何引用setw, setfill等格式化操作符，而不用引用整个std.

Cygwin下安装QT3后，发现qmake找不到，在/usr/lib/qt3/bin下找到了，运行qmake -makefile后，发现它以绝对路径的形式，生成了Makefile,所以运行make后可以找到uic, 等。
安装xserver,需要安装以下包。
xorg-server (required, the Cygwin/X X Server)

> nit (required, scripts for starting the X server: xinit, startx, startwin (and a shortcut on the Start Menu to run it), startxdmcp.bat )

> rg-docs (optional, man pages)

> art-menu-icons (optional, adds icons for X Clients to the Start menu)

You may also select any X client programs you want to use, and any fonts you would like to have available.


http://x.cygwin.com/docs/ug/setup-cygwin-x-installing.html

2010，3，4
Boost的编译其实是相当简单，运行源码目录下的bootstrap以编译bjam。完了之后，使用bjam进行编译，比如：bjam --build-drit=D:tmp install
这样的话，编译产生的临时文件就会放在D:tmp下，编译产生的结果即一个include,和一个lib将会放在C:Boost下。

WTL 安装WTL80，将WTL80解压至c:wtl80目录。将c:wtl80appwizsetup80.js复制为setup90.js，用记事本打开setup90.js，打开编辑菜单中的替换命令，将8.0全部替换为9.0后保存。运行setup90.js即可将WTL80的应用程序向导安装至VS2008中。打末VS2008，选择Tools->Options->Projects and Solutions->VC++ Directories，在include中将c:wtl80include加入。
setup80x.js对应于vs2008 express。。。Wizard是可以安装成功，但是问题是新版的 windows sdk中不包括ATL， 所以无法使用。当然你可以从其它地方拷贝。麻烦，，还是不弄了。。。
Platform SDK是老版的提法。。现在的Windows SDK包括了Platform SDK和.net FrameWork SDK.

ColorTheme使用
首先下载 color-theme.el 1，然后把它放在你的加载路径里面，最后在你的 ~/.emacs 里面加上 (require 'color-theme) 就可以使用了。现在重新启动一 下 Emacs ，然后就可以用 M-x color-theme-select 来选择你喜欢的颜色主题 了，它会打开一个列表，在每个列表项目上回车会应用那个颜色主题，如果选上 了某个主题，按一下 d ，会出现类似：

> lor-theme-sitaramv-nt is an interactive compiled Lisp function in `color-theme.el'. (color-theme-sitaramv-nt) Black foreground on white background. Includes faces for font-lock, widget, custom, speedbar. [back]
的东西，其中那句 (color-theme-sitaramv-nt) 加入到你的 ~/.emacs 里面去， 就可以永久应用这个颜色主题了


太赞了，Chrome 使用SSH(Socks5)代理的方法： 在Chrome 快捷方式属性的目标栏后面加参数--proxy-server=socks5://127.0.0.1:7070 

unified diff format统一差异格式
一种标准的文件比较格式, 不同的行之前标上'+'或者'-'表示不同的文件, 新文件用'+'表示, 旧文件用'-'表示
@@表示不同出现在哪一行

diff -u 可以产生这种格式的补丁文件,它与diff -c命令产生的context diff不一样，后者更适合于大量修改的源代码之间的补丁.
前者的好处在于便于人阅读，而且可以直接patch

emacs高亮前行模式：hl-line-mode
emacs折叠行命令：toggle-truncate-lines

>  having a problem with Qt 3.3.2 on Windows XP using the MS Visual C++.net (2002) compiler. Anytime in my code where I try to select a character from a QString using the [] operator (such as: QChar temp = string[index]), I get the following compiler error: error C2666: 'QString::operator`[]'' : 3 overloads have similar conversions c:Qt3.3.2includeqstring.h(641): could be 'QCharRef QString::operator [](int)' c:Qt3.3.2includeqstring.h(639): or 'QChar QString::operator [](int) const' or 'built-in C++ operator[(const char *, unsigned int)' while trying to match the argument list '(QString, unsigned int)' I don't believe I had this problem with Qt 3.3.1.I had this problem as well and didn't bother to report it to qt-bugs@trolltech.com nor really investigate what was happening. I just worked around the problem.Instead of:QChar c = parse[i];try: QChar c = parse.operator[](i); Note that this problem affects Visual C++ 7.1 (aka 2003) as well. You may wish to follow up with qt-bugs.

.bat中的set设置环境变量时，变量名和内容之间不能有空格，，，，这样是不对的，，set lib =

生成vcproj文件：
```
qmake -project -t vcapp
qmake
```
或者：
```
qmake -project
qmake -tp vc
```
qt3的默认motif风格极其丑陋，yum install qt3-config后，运行qtconfig可进行调整。。
sudo apt-get install libqt3-mt-dev
yum install qt-devel
`proxy server:port number proxy=http://mycache.mydomain.com:3128`
cygcheck
 sudo apt-get install libqt3

ecch $0
chsh -s new_shell user

Q_UINT32 QHostAddress::toIPv4Address () const
Returns the IPv4 address as a number.
For example, if the address is 127.0.0.1, the returned value is 2130706433 (i.e. 0x7f000001).

This value is only valid when isIp4Addr() returns TRUE.

See also toString().

But inet_addr("127.0.0.1") returns 0x1000007f.

174.36.30.71     www.dropbox.com

Tramp:

Tramp 在windows下使用有些麻烦。
首先需要先下载plink.exe 。

Please try PuTTY (or the plink tool that comes with it). There is a
Tramp method "plink", so you only need to add the PuTTY directory to
%PATH% and then C-x C-f /plink:user@host:/some/file RET.
