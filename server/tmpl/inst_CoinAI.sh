#!/bin/bash

Port="{{.Port}}"
UserID="{{.UserID}}"
NowPath=$(pwd)
StartName="CoinAI.net-${Port}"
DirPath="${NowPath}/${StartName}"

rm -rf ${StartName}
mkdir ${StartName}
cd ${StartName} || exit

################## 环境搭建环节 ########################
echo "======== 环境检测 ========"

if [[ $(command -v npm) ]]; then
  echo "检测到已安装 npm , 继续执行"
else
  echo "未安装 npm , 开始安装 nodejs"
  curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - &&
    sudo apt-get install -y nodejs
fi

if [[ $(command -v pm2) ]]; then
  echo "已安装 pm2 , 继续执行"
else
  echo "未安装 pm2 , 开始安装"
  npm install -g pm2
fi

if [[ $(command -v pm2) ]]; then
  echo ""
else
  echo -e "
pm2 安装失败
请手动依次执行以下命令,然后再重新执行该脚本:
\033[32m

  curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - &&
    sudo apt-get install -y nodejs

  npm install -g pm2

\033[0m
"
  exit 1
fi

################ 侦测系统环境 ##########################

SystemType=$(arch)

DownLoadPath="https://raw.githubusercontent.com/AItrade-mo7/CoinAIPackage/main/CoinAI.net_x86_64"
if [[ ${SystemType} =~ "aarch64" ]]; then
  DownLoadPath="https://raw.githubusercontent.com/AItrade-mo7/CoinAIPackage/main/CoinAI.net_aarch64"
fi

################ 启动脚本 ##########################
echo "======== 生成 启动脚本 ========"
StartShellPath="${DirPath}/Reboot.sh"

cat >"${StartShellPath}" <<END
#!/bin/bash

echo "===== 下载可执行文件 ====="

cd "${DirPath}" || exit

pm2 delete "${StartName}"

rm -rf "${StartName}" &&
  curl -o "${StartName}" "${DownLoadPath}"

sudo chmod 777 "${StartName}"

echo "===== 启动服务 ====="

pm2 start "${StartName}" --name "${StartName}"

END

################ 停止脚本 ##########################
echo "======== 生成 停止脚本 ========"
StopShellPath="${DirPath}/Shutdown.sh"

cat >"${StopShellPath}" <<END
#!/bin/bash

pm2 delete "${StartName}"
rm -rf "$(pwd)"

END

################ 配置文件 ##########################
echo "======== 生成 配置文件 ========"
ConfigFilePath="${DirPath}/app_env.json"

cat >"${ConfigFilePath}" <<END
{
  "Port": "${Port}",
  "UserID": "${UserID}"
}
END

sudo chmod -R 777 "$(pwd)"

# echo "======== 启动服务 ========"
source "${StartShellPath}"
