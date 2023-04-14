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
)

func ChangePassword(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("密码修改失败", "设备异常"))
	}

	var json struct {
		Email         string `bson:"Email"`
		Code          string `bson:"Code"`
		Password      string `bson:"Password"`
		AgainPassword string `bson:"AgainPassword"`
	}
	mFiber.Parser(c, &json)

	// 验证邮箱和密码
	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		return c.JSON(result.Fail.With("邮箱格式不正确", json.Email))
	}

	if len(json.Password) < 16 {
		return c.JSON(result.Fail.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	}

	if json.Password != json.AgainPassword {
		return c.JSON(result.Fail.WithMsg("两次密码不一致"))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		Email: json.Email,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	if len(UserDB.UserID) != 32 {
		return c.JSON(result.ErrAccount.WithData("该邮箱尚未注册"))
	}

	// 验证 新邮箱的 验证码
	err = taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: json.Email,
		Code:  json.Code,
	})
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	err = UserDB.ChangePassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	taskPush.DelEmailCode(json.Email)

	return c.JSON(result.Succeed.WithMsg("修改成功"))
}
