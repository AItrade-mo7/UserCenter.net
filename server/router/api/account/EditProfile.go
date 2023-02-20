package account

import (
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type EditUserType struct {
	OldEmailCode string
	NewEmail     string
	NewEmailCode string
	Avatar       string
	NickName     string
	SecurityCode string
	Password     string
}

func EditProfile(c *fiber.Ctx) error {
	var json EditUserType
	mFiber.Parser(c, &json)

	// UserID, err := middle.TokenAuth(c)
	// if err != nil {
	// 	return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	// }
	// UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
	// 	UserID: UserID,
	// })
	// if err != nil {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	// reg, _ := regexp.Compile("[\u4e00-\u9fa5_a-zA-Z0-9_]{2,12}")
	// match := reg.MatchString(json.NickName)
	// if match {
	// } else {
	// 	return c.JSON(result.Fail.WithMsg("昵称不符合规范"))
	// }

	// // 记录老旧的邮箱
	// // oldEmail := UserDB.AccountData.Email

	// Email_edit := len(json.NewEmail) > 1 && json.NewEmail != UserDB.AccountData.Email
	// NickName_edit := json.NickName != UserDB.AccountData.NickName
	// Avatar_edit := len(json.Avatar) > 2 && json.Avatar != UserDB.AccountData.Avatar
	// SecurityCode_edit := len(json.SecurityCode) > 2 && json.SecurityCode != UserDB.AccountData.SecurityCode

	// FK := bson.D{{
	// 	Key:   "UserID",
	// 	Value: UserDB.UserID,
	// }}
	// UK := bson.D{}

	// if NickName_edit {
	// 	UK = append(UK, bson.E{
	// 		Key: "$set",
	// 		Value: bson.D{
	// 			{
	// 				Key:   "NickName",
	// 				Value: json.NickName,
	// 			},
	// 		},
	// 	})
	// }
	// if Avatar_edit {
	// 	UK = append(UK, bson.E{
	// 		Key: "$set",
	// 		Value: bson.D{
	// 			{
	// 				Key:   "Avatar",
	// 				Value: json.Avatar,
	// 			},
	// 		},
	// 	})
	// }

	// // 验证老邮箱验证码
	// if Email_edit || SecurityCode_edit {
	// 	// 密码验证
	// 	err = UserDB.CheckPassword(json.Password)
	// 	if err != nil {
	// 		return c.JSON(result.ErrLogin.WithMsg(err))
	// 	}

	// 	// 验证码验证
	// 	err = verifyCode.CheckEmailCode(verifyCode.CheckEmailCodeParam{
	// 		Email: UserDB.AccountData.Email,
	// 		Code:  json.OldEmailCode,
	// 	})
	// 	if err != nil {
	// 		UserDB.DB.Close()
	// 		return c.JSON(result.ErrEmailCode.WithData(mStr.ToStr(err)))
	// 	}
	// }

	// if SecurityCode_edit {
	// 	UK = append(UK, bson.E{
	// 		Key: "$set",
	// 		Value: bson.D{
	// 			{
	// 				Key:   "SecurityCode",
	// 				Value: json.SecurityCode,
	// 			},
	// 		},
	// 	})
	// }

	// if Email_edit {

	// 	if err != nil {
	// 		UserDB.DB.Close()
	// 		return c.JSON(result.ErrEmailCode.WithData(mStr.ToStr(err)))
	// 	}
	// 	// 验证新邮箱存不存在
	// 	var dbData dbType.AccountTable
	// 	FK := bson.D{{
	// 		Key:   "Email",
	// 		Value: json.NewEmail,
	// 	}}
	// 	UserDB.DB.Table.FindOne(UserDB.DB.Ctx, FK).Decode(&dbData)

	// 	if len(dbData.UserID) == 32 {
	// 		UserDB.DB.Close()
	// 		return c.JSON(result.Fail.WithMsg("新邮箱已存在"))
	// 	}

	// 	UK = append(UK, bson.E{
	// 		Key: "$set",
	// 		Value: bson.D{
	// 			{
	// 				Key:   "Email",
	// 				Value: json.NewEmail,
	// 			},
	// 		},
	// 	})
	// }

	// if len(UK) < 1 {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.Fail.WithMsg("未作出任何更改"))
	// }

	// UK = append(UK, bson.E{
	// 	Key: "$set",
	// 	Value: bson.D{
	// 		{
	// 			Key:   "UpdateTime",
	// 			Value: mTime.GetUnixInt64(),
	// 		},
	// 	},
	// })

	// _, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	// if err != nil {
	// 	errStr := fmt.Errorf("数据库更新失败 %+v", err)
	// 	global.LogErr(errStr)
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(errStr)))
	// }
	// UserDB.Update()

	// if Email_edit {
	// message := mStr.Join(
	// 	"您刚刚修改了登录邮箱! 操作设备: ", c.Get("User-Agent"), "\n",
	// 	"新的邮箱为: ", UserDB.AccountData.Email,
	// )

	// go global.Email(global.EmailOpt{
	// 	To: []string{
	// 		UserDB.AccountData.Email,
	// 		oldEmail,
	// 	},
	// 	Subject:  "邮箱修改提醒",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message:      message,
	// 		SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
	// 		SecurityCode: UserDB.AccountData.SecurityCode,
	// 	},
	// }).Send()
	// }

	// if SecurityCode_edit {
	// message := mStr.Join(
	// 	"您刚刚修改了防伪标识符! 操作设备: ", c.Get("User-Agent"), "\n",
	// 	"新的防伪标识符为: ", UserDB.AccountData.SecurityCode,
	// )

	// go global.Email(global.EmailOpt{
	// 	To: []string{
	// 		UserDB.AccountData.Email,
	// 		oldEmail,
	// 	},
	// 	Subject:  "防伪标识符修改提醒",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message:      message,
	// 		SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
	// 		SecurityCode: UserDB.AccountData.SecurityCode,
	// 	},
	// }).Send()
	// }

	// var userinfoData apiType.UserInfo
	// jsonStr := mJson.ToJson(UserDB.AccountData)
	// jsoniter.Unmarshal(jsonStr, &userinfoData)

	// UserDB.DB.Close()
	return c.JSON(result.Succeed.WithMsg("修改成功"))
}
