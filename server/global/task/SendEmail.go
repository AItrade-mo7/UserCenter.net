package task

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

func TestSendMsg() {
	Cont := mTask.ToMapData(mTask.SendEmail{
		From: "AITrade",
		To: []string{
			"mo7@mo7.cc",
		},
		Subject:  "Subjectxxx",
		TmplName: "SysEmail",
		SendData: mTask.SysEmailParam{
			Title:        "系统提示",
			Message:      "启动系统:",
			Content:      "这是一封来自系统的测试邮件",
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			Source:       Source,
			SecurityCode: "trade.mo7.cc",
		},
	})

	err := NewTask(NewTaskOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: "测试一波发送",
	})

	fmt.Println(err)
}
