package ready

import (
	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
)

func StartEmail() {
	err := taskPush.SysEmail(taskPush.SysEmailOpt{
		From:        config.SysName,
		Subject:     "系统启动",
		Title:       config.SysName + " 系统启动",
		Message:     "系统启动",
		Content:     mJson.Format(config.AppInfo),
		Description: "系统启动邮件",
	})
	global.Run.Println("系统启动邮件已发送", err)
}
