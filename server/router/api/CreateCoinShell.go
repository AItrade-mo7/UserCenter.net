package api

import (
	"DataCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type CreateCoinShellParam struct {
	Port   string `bson:"Port"`
	UserID string `bson:"UserID"` // 用户 ID
}

func CreateCoinShell(c *fiber.Ctx) error {
	var json CreateCoinShellParam
	mFiber.Parser(c, &json)

	if len(json.Port) > 2 {
		return c.JSON(result.Fail.WithData("请填写端口号"))
	}

	return c.JSON(result.Succeed.WithData("尚在开发中"))
}
