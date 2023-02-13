package taskPush

import (
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

// === 系统邮件 ====
type SysEmailOpt struct {
	From         string   // 缺省 AItrade
	To           []string // 缺省 trade@mo7.cc
	Subject      string
	Title        string
	Message      string
	Content      string
	Description  string
	SecurityCode string
}

func SysEmail(opt SysEmailOpt) error {
	if len(opt.SecurityCode) < 2 {
		opt.SecurityCode = "trade.mo7.cc"
	}

	if len(opt.From) < 2 {
		opt.From = "AItrade"
	}

	Cont := mTask.ToMapData(mTask.SendEmail{
		From:     opt.From,
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
	err := New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: opt.Description,
	})

	return err
}

// === 发送验证码 ====
type CodeEmailOpt struct {
	To           []string
	VerifyCode   string
	Action       string
	SecurityCode string
}

func CodeEmail(opt CodeEmailOpt) error {
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
	err := New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: "验证码邮件",
	})

	return err
}

// 注册成功通知

type RegisterEmailOpt struct {
	To       []string
	Password string
}

func RegisterEmail(opt RegisterEmailOpt) error {
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
	err := New(NewOpt{
		TaskType:    "SendEmail",
		Content:     Cont,
		Description: "注册成功邮件通知",
	})
	return err
}
