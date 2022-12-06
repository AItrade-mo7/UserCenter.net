package global

import (
	"time"

	"DataCenter.net/server/global/config"
	"DataCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 初始化目录列表
	config.DirInit()

	// 初始化日志系统 保证日志可用
	mCycle.New(mCycle.Opt{
		Func:      LogInit,
		SleepTime: time.Hour * 24,
	}).Start()

	ServerEnvInit()

	go Email(EmailOpt{
		To:       config.Email.To,
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "系统初始化完成",
			SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
		},
	}).Send()
	Log.Println(`系统初始化完成`)
}
