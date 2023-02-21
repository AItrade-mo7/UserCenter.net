package account

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mFetch"
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

	UserAgent := config.SysName
	Path := "/api/await/SendEmailCode"
	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: config.SysEnv.MessageBaseUrl,
		Path:   Path,
		Data:   mJson.ToJson(json),
		Header: map[string]string{
			"Auth-Encrypt": config.ClientEncrypt(Path + UserAgent),
			"User-Agent":   UserAgent,
		},
	})
	resData, err := fetch.Post()
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
