package coinAI

import (
	"DataCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithMsg("列表接口尚在开发中"))
}
