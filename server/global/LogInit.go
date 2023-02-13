package global

import (
	"fmt"
	"log"

	"UserCenter.net/server/global/config"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

var (
	Log *log.Logger // 系统日志& 重大错误或者事件
	Run *log.Logger // 运行日志
)

func LogInit() {
	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	Run = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Run",
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
	str := fmt.Sprintf("系统错误: %+v", sum)
	Log.Println(str)

	message := ""
	if len(sum) > 0 {
		message = mStr.ToStr(sum[0])
	}
	content := mJson.Format(sum)

	go taskPush.SysEmail(taskPush.SysEmailOpt{
		From:        taskPush.Source,
		Subject:     "系统错误",
		Title:       taskPush.Source + " 系统出错",
		Message:     message,
		Content:     content,
		Description: "出现系统错误",
	})
}
