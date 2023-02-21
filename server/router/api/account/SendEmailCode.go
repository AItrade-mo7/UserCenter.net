package account

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mRes"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type SendEmailCodeParam struct {
	Email  string
	Action string
}

func SendEmailCode(c *fiber.Ctx) error {
	var json SendEmailCodeParam
	mFiber.Parser(c, &json)

	resData, err := taskPush.Request(taskPush.RequestOpt{
		Origin: config.SysEnv.MessageBaseUrl,
		Path:   "/api/await/SendEmailCode",
		Data:   mJson.ToJson(json),
	})
	if err != nil {
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	var resObj mRes.ResType
	jsoniter.Unmarshal(resData, &resObj)

	if resObj.Code < 0 {
		return c.JSON(result.ErrEmailCode.WithMsg(resObj.Msg))
	}

	return c.JSON(result.Succeed.WithMsg("验证码已发送"))
}
