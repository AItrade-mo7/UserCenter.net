package main

import (
	_ "embed"

	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/ready"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mEncrypt"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	// 初始化系统参数
	global.Start()

	// 数据准备
	ready.Start()

	// 启动 http 监听服务
	// router.Start()

	taskPush.CodeEmail(taskPush.CodeEmailOpt{
		To:         []string{"meichangliang@outlook.com"},
		VerifyCode: mEncrypt.GetUUID(),
		Action:     "测试",
	})
}
