---
title: "Gitlab Merge Request 的三个选项"
date: 2021-01-14T21:34:39+08:00
draft: false
tags: [git]
---

gitlab 仓库的 merge request 有三个配置选项：

![Image](gitlab-merge-request.assets/Image.png) 



+ 选项 1 (merge commit)，只有源分支会被测试，因此 master 上有被破坏的风险。
+ 选项 2 (Merge commit with semi-linear history)，除了测试源分支外，还会测试合并后的节点，因此 master 分支会比较安全，但是会浪费比较多的测试时间，同时开发人员需要进行 rebase。
+ 选项 3 (Fast-Forward merge) 的测试策略和第一第二是一样的

选项 2 和选项 3 只是生成的 commit history 会不同

以下的三个图，分别对应上面的三个选项生成的 commit history

![image-20210114221859780](gitlab-merge-request.assets/image-20210114221859780.png)

