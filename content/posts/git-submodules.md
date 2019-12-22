---
title: "Git Submodule 教程"
date: 2019-12-22T14:32:54+08:00
draft: false
---

## 为现有的仓库增加 submodule

`git submodule add (repository) (directory)`
在执行了以上命令后，在 git 仓库的根目录下，会出现一个 `.gitmodules`，里面会记录当前添加的作为 `submodule` 的 git 仓库的地址。

提交这个 `.gitmodules` 文件：
```
git add .gitmodules
git commit -m"add sumodule file"
git push
```

经过以上的步骤，我们就把为我们的 git 仓库添加了一个 `submodule`


## clone 现有的包含 submodule 的仓库

重新`clone` 我们的仓库，会发现，里面只有一个空的目录。这是因为，默认情况下，`git clone` 不会把 `submodule` 同时 `clone` 下来。为了把`submodule` 也 `clone` 下来，需要运行一下的 `submodule` 指令

```
git submodule init
git submodule update
// 也有可能 sumbodule 可能会包含其他的 git repo, 所以可以使用一下的命令
git submodule update --init --recursive

```

`submodule` 也是一个普通的 repo，我们可以在目录中，执行`log`， `status`， `checkout`和`branch`


## 删除 submodule 

删除有点繁琐：

1. 首先删除对应的 submodule 目录 `rm -fr submodule`
2. 删除 `.gitmodules` 中的相关条目
3. 删除 `.git/config` 中相关的条目
4. `git rm --cached <path/to/submodule>`































