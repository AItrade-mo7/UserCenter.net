package global

import (
	"fmt"
	"log"

	"UserCenter.net/server/global/config"
	"UserCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mTime"
)

var Log *log.Logger // 系统日志& 重大错误或者事件
func LogInit() {
	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	// 设定清除log
	mLog.Clear(mLog.ClearParam{
		Path:      config.Dir.Log,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})

	// 将方法重载到 config 内部,便于执行
	config.LogErr = LogErr
	config.Log = Log
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误 : %+v", sum)
	Email := Email(EmailOpt{
		To:       config.Email.To,
		Subject:  "LogErr",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: str,
			SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
		},
	})
	Log.Println(str)
	go Email.Send()
}
