---
title: "Vuetify Icon"
date: 2019-12-07T20:50:35+08:00
draft: false
tags: [JavaScript, CSS]
---

安装 Vuetify 的步骤

1. `yarn global add @vue/cli`
2. `vue create my-app`
3. `vue add vuetify`

第三步后， 会自动再 `public/index.html` 文件下添加字体，是默认的 `material design icon`

因此再安装后，使用 `v-icon` 时，需要使用 `mdi` 前缀。比如： `<v-icon>mdi-basketball</v-icon>`


但是再使用一些文档中的例子的时候，比如从 codepen 中粘贴过来是，里面用的字体可能是不带 `mdi` 前缀的，因此需要按如下的步骤安装 `material icon`

```
yarn add material-design-icons-iconfont -D
```

并在`main.js` ， `index.js` 或者 `app.js` 中，添加如下：

```
import 'material-design-icons-iconfont/dist/material-design-icons.css' 
```


Material design icon 和  Material icon 是不同的两套图标