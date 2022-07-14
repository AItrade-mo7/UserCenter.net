package account

import (
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"DataCenter.net/server/utils/verifyCode"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

type RegisterParam struct {
	Email string `bson:"Email"`
	Code  string `bson:"Code"`
}

func Register(c *fiber.Ctx) error {
	var json RegisterParam
	mFiber.Parser(c, &json)

	err := verifyCode.CheckEmailCode(verifyCode.CheckEmailCodeParam{
		Email: json.Email,
		Code:  json.Code,
	})
	if err != nil {
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		Email: json.Email,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	if len(UserDB.UserID) == 32 {
		UserDB.DB.Close()
		return c.JSON(result.ErrAccountRepeat.WithData("该邮箱已注册"))
	}

	err = UserDB.Register(json.Email)
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	UserDB.DB.Close()
	return c.JSON(result.Succeed.With("注册成功", "密码已发送至您的邮箱，请注意查收"))
}
