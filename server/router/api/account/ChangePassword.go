package account

import (
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func ChangePassword(c *fiber.Ctx) error {
	// var json struct {
	// 	Email         string `bson:"Email"`
	// 	Code          string `bson:"Code"`
	// 	Password      string `bson:"Password"`
	// 	AgainPassword string `bson:"AgainPassword"`
	// }
	// mFiber.Parser(c, &json)

	// // 验证邮箱和密码
	// isEmail := mVerify.IsEmail(json.Email)
	// if !isEmail {
	// 	return c.JSON(result.Fail.With("邮箱格式不正确", json.Email))
	// }

	// if len(json.Password) != 32 {
	// 	return c.JSON(result.Fail.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	// }

	// if json.Password != json.AgainPassword {
	// 	return c.JSON(result.Fail.WithMsg("两次密码不一致"))
	// }

	// UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
	// 	Email: json.Email,
	// })
	// if err != nil {
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	// if len(UserDB.UserID) != 32 {
	// 	return c.JSON(result.ErrAccount.WithData("该邮箱尚未注册"))
	// }

	// err = UserDB.ChangePassword(json.Password)
	// if err != nil {
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	return c.JSON(result.Succeed.WithMsg("修改成功"))
}
