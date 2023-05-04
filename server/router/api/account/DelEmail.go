package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// 设置主要的 Email
type DelEmailParam struct {
	Email     string // 要删除的邮箱
	EmailCode string // 邮箱的验证码
	Password  string // 账户的密码
}

func DelEmail(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("设置Email失败", "设备异常"))
	}

	var json SetMainEmailParam
	mFiber.Parser(c, &json)

	// 输入参数格式验证
	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		return c.JSON(result.Fail.With("邮箱格式不正确", json.Email))
	}
	if len(json.Password) < 16 {
		return c.JSON(result.Fail.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	}

	// 当前登录账户信息
	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: userID,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()

	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}
	// 验证验证码
	err = taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: UserDB.Data.Email,
		Code:  json.EmailCode,
	})
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	// 如果它是主要
	if json.Email == UserDB.Data.Email {
		return c.JSON(result.Fail.WithMsg("不可删除主要邮箱"))
	}

	// 验证新邮箱是否为 List 一员
	isUserEmail := false
	NewEmailList := []string{}
	for _, val := range UserDB.Data.UserEmail {
		if val == json.Email {
			isUserEmail = true
		} else {
			NewEmailList = append(NewEmailList, val)
		}
	}

	if !isUserEmail {
		return c.JSON(result.Fail.WithMsg("当前邮箱不在列表当中"))
	}

	taskPush.SysEmail(taskPush.SysEmailOpt{
		To:             UserDB.Data.UserEmail,
		Subject:        "登录提醒",
		Title:          "你正在删除邮箱地址！",
		Message:        "系统侦测到您正在删除如下邮箱地址:",
		Content:        json.Email,
		Description:    "删除邮件通知",
		EntrapmentCode: UserDB.Data.EntrapmentCode,
	})

	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.Data.UserID,
	}}
	UK := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "UserEmail",
				Value: NewEmailList,
			},
		},
	}}
	UK = append(UK, bson.E{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "UpdateTime",
				Value: mTime.GetUnixInt64(),
			},
		},
	})

	_, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	taskPush.DelEmailCode(UserDB.Data.Email)
	return c.JSON(result.Succeed.WithData("删除成功"))
}
