---
title: "Maven 中使用本地 jar 包"
date: 2019-08-20T20:48:32+08:00
draft: false
tags: [Java]
---

你可以包含本地的 jar, 像下面一样:

![image-20191217004922024](mvn-use-local-jar.assets/image-20191217004922024.png)


但是, 在使用 `mvn package` 时,这个 jar 不会被打包进去,除非使用下面的方式`安装`

```
mvn install:install-file 
-Dfile="C:\Users\xxx\lib\LoginRadius-1.0.jar" -DgroupId=LoginRadius -DartifactId=LoginRadius -Dversion=1.0 -Dpackaging=jar`
```

https://stackoverflow.com/questions/19065666/how-to-include-system-dependencies-in-war-built-using-maven