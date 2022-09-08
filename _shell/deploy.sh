#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############
npm install -g pnpm
npm install -g pm2

pnpm install
echo "开始打包" &&
  pnpm run build

echo "停止 pm2 服务" &&
  pm2 delete ${startName}

echo "移动文件到 ProdProject 目录"
cp -r ${outPutPath}"/." ${deployPath}"/"

cd ${deployPath}

echo "启动 pm2 服务"
pm2 start ./${buildName} --name ${startName}
