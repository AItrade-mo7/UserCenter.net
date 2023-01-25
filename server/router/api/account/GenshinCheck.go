package account

import (
	"DataCenter.net/server/genshin"
	"DataCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

type GenshinCheckParam struct {
	Cookie string
}

func GenshinCheck(c *fiber.Ctx) error {
	var json GenshinCheckParam
	mFiber.Parser(c, &json)

	if len(json.Cookie) < 12 {
		return c.JSON(result.Fail.WithMsg("Cookie 长度不足"))
	}

	resData, resErr := genshin.SignIn(json.Cookie)

	if resErr != nil {
		return c.JSON(result.Fail.WithMsg("出现了未知错误,请截图反馈给开发者" + mStr.ToStr(resErr)))
	}

	return c.JSON(result.Succeed.WithMsg(resData))
}
