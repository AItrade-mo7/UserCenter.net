package account

import (
	"fmt"

	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"DataCenter.net/server/utils/verifyCode"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
)

func SendEmailCode(c *fiber.Ctx) error {
	var json verifyCode.SendCodeParam
	mFiber.Parser(c, &json)

	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		emailErr := fmt.Errorf("email 格式不正确 %+v", json.Email)
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		Email: json.Email,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	if len(UserDB.UserID) == 32 {
		json.SecurityCode = UserDB.AccountData.SecurityCode
	}

	err = verifyCode.CheckAndSendCode(json)
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrEmail.WithMsg(err))
	}

	UserDB.DB.Close()
	return c.JSON(result.Succeed.WithMsg("验证码已发送"))
}
