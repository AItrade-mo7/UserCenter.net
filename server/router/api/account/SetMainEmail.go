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
type SetMainEmailParam struct {
	Email     string
	EmailCode string
	Password  string
}

func SetMainEmail(c *fiber.Ctx) error {
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
	if len(json.Password) != 32 {
		return c.JSON(result.Fail.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	}

	// 验证验证码
	err := taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: json.Email,
		Code:  json.EmailCode,
	})
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
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
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()
	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}
	// 验证新邮箱是否为 List 一员

	isUserEmail := false
	for _, val := range UserDB.Data.UserEmail {
		if val == json.Email {
			isUserEmail = true
			break
		}
	}

	if !isUserEmail {
		return c.JSON(result.Fail.WithMsg("当前邮箱未添加！"))
	}

	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.Data.UserID,
	}}
	UK := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "Email",
				Value: json.Email,
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

	return c.JSON(result.Succeed.WithMsg("设置主要的Email"))
}
