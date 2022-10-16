package coinAI

import (
	"DataCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func Remove(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithMsg("删除接口尚在开发中"))
}
