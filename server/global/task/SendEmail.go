package task

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

func TestSendMsg() {
	Cont := mTask.ToMapData(mTask.SendEmail{
		From: "AITrade.net",
		To: []string{
			"meichangliang@outlook.com",
			"mo7@mo7.cc",
		},
		Subject:  "测试邮件!",
		TmplName: "SysEmail",
		SendData: mTask.SysEmailParam{
			Title:        "测试邮件",
			Message:      "测试一下邮件的发送",
			Content:      "这是一封来自系统的测试邮件",
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			Source:       Source,
			SecurityCode: "trade.mo7.cc",
		},

		// TmplName: "CodeEmail",
		// SendData: mTask.CodeEmailParam{
		// 	VerifyCode:   "045685",
		// 	Action:       "测试",
		// 	SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
		// 	Source:       Source,
		// 	SecurityCode: "trade.mo7.cc",
		// },

		// TmplName: "RegisterEmail",
		// SendData: mTask.RegisterParam{
		// 	Password:     "SDGASDas",
		// 	SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
		// 	Source:       Source,
		// 	SecurityCode: "trade.mo7.cc",
		// },
	})

	start := mTime.GetUnixInt64()
	fmt.Println("开始创建 5 个任务", start)
	for i := 0; i < 5; i++ {
		err := NewTask(NewTaskOpt{
			TaskType:    "SendEmail",
			Content:     Cont,
			Description: "测试一波发送",
		})
		fmt.Println(i, err)
	}

	end := mTime.GetUnixInt64()
	fmt.Println("end", end)
	fmt.Println("用时", end-start)
}
