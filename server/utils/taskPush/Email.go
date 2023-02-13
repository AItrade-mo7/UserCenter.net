package taskPush

import (
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

// === 系统邮件 ====
type SysEmailOpt struct {
	To           []string
	Subject      string
	Title        string
	Message      string
	Content      string
	Description  string
	SecurityCode string
}

func SysEmail(opt SysEmailOpt) {
	if len(opt.SecurityCode) < 2 {
		opt.SecurityCode = "trade.mo7.cc"
	}

	Cont := mTask.ToMapData(mTask.SendEmail{
		From:     "AItrade",
		To:       opt.To,
		Subject:  opt.Subject,
		TmplName: "SysEmail",
		SendData: mTask.SysEmailParam{
			Title:        opt.Title,
			Message:      opt.Message,
			Content:      opt.Content,
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			Source:       Source,
			SecurityCode: opt.SecurityCode,
		},
	})
	New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: opt.Description,
	})
}

// === 发送验证码 ====
type CodeEmailOpt struct {
	To           []string
	VerifyCode   string
	Action       string
	SecurityCode string
}

func CodeEmail(opt CodeEmailOpt) {
	Cont := mTask.ToMapData(mTask.SendEmail{
		From:     "AItrade",
		To:       opt.To,
		Subject:  "请查收您的验证码",
		TmplName: "CodeEmail",
		SendData: mTask.CodeEmailParam{
			VerifyCode:   opt.VerifyCode,
			Action:       opt.Action,
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			Source:       Source,
			SecurityCode: opt.SecurityCode,
		},
	})
	New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: "验证码邮件",
	})
}

// 注册成功通知

type RegisterEmailOpt struct {
	To       []string
	Password string
}

func RegisterEmail(opt RegisterEmailOpt) {
	Cont := mTask.ToMapData(mTask.SendEmail{
		From:     "AItrade",
		To:       opt.To,
		Subject:  "注册成功！",
		TmplName: "RegisterEmail",
		SendData: mTask.RegisterParam{
			Password:     opt.Password,
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			Source:       Source,
			SecurityCode: "trade.mo7.cc",
		},
	})
	New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: "注册成功邮件通知",
	})
}
