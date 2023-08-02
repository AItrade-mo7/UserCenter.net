package account

import (
	"fmt"

	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/taskPush"
	"UserCenter.net/sysPublic/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
)

type RegisterParam struct {
	Email          string `bson:"Email"`
	Code           string `bson:"Code"`
	EntrapmentCode string `bson:"EntrapmentCode"` // 防钓鱼码
}

func Register(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("注册失败", "设备异常"))
	}

	var json RegisterParam
	mFiber.Parser(c, &json)

	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		err := fmt.Errorf("邮箱格式不正确 %+v", json.Email)
		return c.JSON(result.ErrEmail.WithMsg(err))
	}

	if len(json.Code) < 1 {
		err := fmt.Errorf("验证码不能为空")
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	if len(json.EntrapmentCode) < 1 {
		emailErr := fmt.Errorf("防钓鱼码不能为空")
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	if len([]rune(json.EntrapmentCode)) > 24 {
		emailErr := fmt.Errorf("防钓鱼码不能大于24位")
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	// 在这里检查验证码
	err := taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
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
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()

	if len(UserDB.UserID) > 0 {
		return c.JSON(result.ErrAccountRepeat.WithData("该邮箱已注册"))
	}

	err = UserDB.Register(dbUser.RegisterOpt{
		Email:          json.Email,
		EntrapmentCode: json.EntrapmentCode,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	taskPush.DelEmailCode(json.Email)

	return c.JSON(result.Succeed.With("注册成功", "密码已发送至您的邮箱，请注意查收"))
}
