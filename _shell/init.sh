#!/bin/bash


function GitSet {
  echo "设置大小写敏感,git忽略权限变更,更改权限"
  git config core.ignorecase false

  git config --global core.fileMode false
  git config core.filemode false

  chmod -R 777 ./
}

## 存储变量

# 项目根目录
path=$(pwd)

# 项目的名字和编译时的名字
startName=${path##*/}
buildName=${startName}

# 最终的输出目录
outPutPath=${path}"/dist"

# 部署目录
deployPath="/root/ProdProject/DataCenter"


