---
title: "JavaScript 立即执行函数表达式 IIFE"
date: 2019-12-29T09:26:02+08:00
draft: false
tags: [JavaScript]
---

IIFE 是指： `JavaScript Immediately-invoked Function Expressions`。 也就是定义一个匿名函数，并且立刻调用它。


比如：
```
function() {
	console.log("Hello World");
}()
```

IFFE 在我们想缩小变量的作用域的时候非常有用，比如如下代码用来显示 Webpack 的打包进度，bar1 作为全局变量存在

```JavaScript
var bar1 = null;
const handler = (percentage, message, ...args) => {

  if (bar1 === null) {
    const cliProgress = require('cli-progress');
    bar1 = new cliProgress.SingleBar({}, cliProgress.Presets.shades_classic);
    bar1.start(100, 0);
  }
  bar1.update(percentage * 100);
  if (percentage * 100 >= 100) {
    bar1.stop();
  }

};
```

使用 `IIFE` 可以改造为：

```javascript
const handler = (() => {
  let progressBar = null;
  return (percentage, message, ...args) => {
    if (progressBar === null) {
      const cliProgress = require('cli-progress');
      progressBar = new cliProgress.SingleBar({}, cliProgress.Presets.shades_classic);
      progressBar.start(100, 0);
    }
    // e.g. Output each progress message directly to the console:
    progressBar.update(percentage * 100);
    if (percentage * 100 >= 100) {
      progressBar.stop();
    }
  }
})();
```

需要注意的是，使用 `arrow function` IIFE 需要使用括号把箭头函数括起来

```
(()=>{
  // do something
})()
```



