package dbUser

import (
	"fmt"
	"strings"

	"UserCenter.net/server/global"
	"UserCenter.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mTime"
)

type RegisterOpt struct {
	Email        string
	SecurityCode string
}

func (dbObj *AccountType) Register(opt RegisterOpt) (resErr error) {
	resErr = nil

	// 检查数据库连接状态
	db := dbObj.DB
	defer db.Close()
	err := db.Ping()
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("注册用户,数据库连接错误 %+v", err)
		global.LogErr(resErr)
		return
	}

	if len(dbObj.Data.UserID) > 10 {
		resErr = fmt.Errorf("该账号已注册，请直接登录")
		return
	}

	newPwd := mEncrypt.RandStr(8) // 生成密码
	UserEmail := []string{}
	UserEmail = append(UserEmail, opt.Email)

	var Body dbType.UserTable
	Body.UserID = mEncrypt.GetUUID()                 // 生成 UserID
	Body.Email = opt.Email                           // 插入邮箱
	Body.UserEmail = UserEmail                       // 插入邮箱
	Body.Avatar = "//file.mo7.cc/AItrade/avatar.png" // 生成默认头像
	Body.NickName = "AItrade用户"                      // 生成昵称,昵称应该为邮箱前缀
	Body.CreateTime = mTime.GetTime().TimeUnix       // 生成创建时间
	Body.UpdateTime = mTime.GetTime().TimeUnix       // 生成更新时间
	Body.SecurityCode = opt.SecurityCode             // 防伪标识符
	Body.Password = mEncrypt.MD5(newPwd)             // 密码加密存储

	str_arr := strings.Split(email, `@`)
	if len(str_arr) > 0 {
		Body.NickName = str_arr[0]
	}

	// 1. 发送邮件告知密码
	err = SendPwd(SendPwdType{
		Email:        Body.Email,
		Password:     newPwd,
		SecurityCode: Body.SecurityCode,
	})

	if err != nil {
		db.Close()
		resErr = err
		return
	}

	// 2. 插入数据库
	_, err = db.Table.InsertOne(db.Ctx, Body)
	if err != nil {
		resErr = fmt.Errorf("注册,插入数据失败 %+v", err)
		global.LogErr(resErr)
		db.Close()
		return
	}

	dbObj.Update()

	return
}

type SendPwdType struct {
	Email        string
	Password     string
	SecurityCode string
}

func SendPwd(opt SendPwdType) error {
	// Email := global.Email(global.EmailOpt{
	// 	To: []string{
	// 		opt.Email,
	// 	},
	// 	Subject:  "注册成功",
	// 	Template: tmpl.RegisterSucceedEmail,
	// 	SendData: tmpl.RegisterSucceedParam{
	// 		SysTime:      mTime.UnixFormat(""),
	// 		Password:     opt.Password,
	// 		SecurityCode: opt.SecurityCode,
	// 	},
	// })

	// return Email.Send()
	return nil
}
