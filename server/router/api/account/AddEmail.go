package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// 新增 Email
type AddEmailParam struct {
	Email     string
	EmailCode string
	Password  string
}

func AddEmail(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("新增Email失败", "设备异常"))
	}

	var json AddEmailParam
	mFiber.Parser(c, &json)

	// 输入参数格式验证
	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		return c.JSON(result.Fail.With("邮箱格式不正确", json.Email))
	}
	if len(json.Password) != 32 {
		return c.JSON(result.Fail.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	}

	// 验证 新邮箱的 验证码
	err := taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: json.Email,
		Code:  json.EmailCode,
	})
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	// 验证新邮箱是否存在
	NewEmailDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		Email: json.Email,
	})
	if err != nil {
		NewEmailDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer NewEmailDB.DB.Close()
	if len(NewEmailDB.UserID) > 0 {
		return c.JSON(result.Fail.WithMsg("当前邮箱已被使用"))
	}
	NewEmailDB.DB.Close()

	// 当前登录账户信息
	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: userID,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()
	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	// 可以新增邮箱了
	UserDB.Data.UserEmail = append(UserDB.Data.UserEmail, json.Email)
	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.Data.UserID,
	}}
	UK := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "UserEmail",
				Value: UserDB.Data.UserEmail,
			},
		},
	}}
	_, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithMsg("Email已新增"))
}
