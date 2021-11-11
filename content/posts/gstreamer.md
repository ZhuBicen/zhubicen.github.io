---
title: "小试 gstreamer"
date: 2021-10-23T22:11:31+08:00
draft: false
tags: [gstreamer]
---

原始的想法是通过 gstreamer 产生RTP/RTCP。在网上找到这么个公式

```

gst-launch videotestsrc ! ffenc_mpeg4 ! rtpmp4vpay send-config=true ! udpsink host=127.0.0.1 port=5000

```

因此想安装试一下，在各个系统下试了好几次，总是有问题，要么是找不到 component，要不是提示别的错误。好像总是安装不正确或者不完全

经过摸索有几个小坑，一个要注意的是 gstreamer 有两个大的版本，一个是 gstreamer 0.1，一个是 gstreamer 1.0，是两个完全不同的版本，可能会有不同的模块，因此搜到的命令如果是使用的 0.1 ，拿到 1.0 的 gstreamer 上就会出现问题

另一个是安装的问题，特别是 MAC 的安装的官方文档说明有问题。它明确说，不支持 MacOS 10.11 (El Capitan) 以上的版本。事实证明，在 M1 版本的 Mac 上没有问题，使用如下命令安装：

```

brew install gstreamer gst-plugins-base gst-plugins-good gst-plugins-bad gst-plugins-ugly gst-libav

```

另外，对于那像魔法一样的 gst-launch 命令，要使用前大概明白他的语法，推荐 Youtube 上的[视频](https://www.youtube.com/watch?v=_yU1kfcC6rY&t=2607s)。最关键的一点是弄清楚他的语法，还有会使用 gst-inspect-1.0 来查看每个组件的能力。比如查看 `rtpmp4vpay` 相关的信息：`gst-launch-1.0 rtpmp4vpay`

在经过了一番折腾后，转化为以下的 pipeline (生成 MP4 RTP stream)，就可以使用了：

```

gst-launch-1.0 videotestsrc ! avenc_mpeg4 ! rtpmp4vpay config-interval=1 ! udpsink host=127.0.0.1 port=5000

```

sdp

```

v=0

m=video 5000 RTP/AVP 96

c=IN IP4 127.0.0.1

a=rtpmap:96 MP4V-ES/90000

```

那么如何转化为 H264 编码的呢，使用 gst-inspect | grep h264 找到对应的 H264 编码就可以了

```

gst-launch-1.0 videotestsrc ! avenc_h264 ! rtph264pay ! udpsink host=127.0.0.1 port=5000

```

对应的 sdp 为，使用 vlc sdp 就可以了

```

v=0

m=video 5000 RTP/AVP 96

c=IN IP4 127.0.0.1

a=rtpmap:96 H264/90000

```

下面两个命令可以分别以屏幕和摄像头为源(在 MAC 上验证通过）

使用 `gst-launch-1.0 avfvideosrc capture-screen=true ! autovideosink` 可以捕获屏幕

参考

[Implementing GStreamer Webcam(USB & Internal) Streaming](https://medium.com/lifesjourneythroughalens/implementing-gstreamer-webcam-usb-internal-streaming-mac-c-clion-76de0fdb8b34)

[GStreamer rtp stream to vlc](https://stackoverflow.com/questions/13154983/gstreamer-rtp-stream-to-vlc)
