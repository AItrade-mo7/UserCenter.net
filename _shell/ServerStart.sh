#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############
sysPublicPath=${sysPublicPath}
NowPath=${NowPath}

echo " =========== 同步公共模块  =========== "

cd "${sysPublicPath}" &&
  git pull &&
  cd "${NowPath}" || exit

echo " =========== 清理目录 =========== "
rm -rf ./logs
rm -rf ./jsonData

echo "整理 mod"
go mod tidy

echo " ========== 开始运行 ========== "
go run main.go
