package ready

import (
	"time"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/config"
	"DataCenter.net/server/tmpl"
)

func Start() {
	if config.SysEnv.RunMod == 0 {
		// 发送邮件
		go global.Email(global.EmailOpt{
			To: []string{
				"meichangliang@mo7.cc",
			},
			Subject:  "ServeStart",
			Template: tmpl.SysEmail,
			SendData: tmpl.SysParam{
				Message: "服务启动",
				SysTime: time.Now(),
			},
		}).Send()
	}
}
