package ready

import (
	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
)

func StartEmail() {
	err := taskPush.SysEmail(taskPush.SysEmailOpt{
		From:        taskPush.Source,
		Subject:     "系统启动",
		Title:       taskPush.Source + " 系统启动",
		Message:     "系统启动",
		Content:     mJson.Format(config.AppInfo),
		Description: "系统启动邮件",
	})
	global.Run.Println("系统启动邮件已发送", err)
}
