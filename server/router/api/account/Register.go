package account

import (
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type RegisterParam struct {
	Email string `bson:"Email"`
	Code  string `bson:"Code"`
}

func Register(c *fiber.Ctx) error {
	var json RegisterParam
	mFiber.Parser(c, &json)

	// 在这里检查验证码
	err := taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: json.Email,
		Code:  json.Code,
	})
	if err != nil {
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	// UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
	// 	Email: json.Email,
	// })
	// if err != nil {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	// if len(UserDB.UserID) == 32 {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrAccountRepeat.WithData("该邮箱已注册"))
	// }

	// err = UserDB.Register(json.Email)
	// if err != nil {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	// UserDB.DB.Close()
	return c.JSON(result.ErrDB.With("注册成功", "密码已发送至您的邮箱，请注意查收"))
}
