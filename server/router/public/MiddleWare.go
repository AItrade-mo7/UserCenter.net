package public

import (
	"UserCenter.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", config.SysName+"/api/public")

	return c.Next()
}
