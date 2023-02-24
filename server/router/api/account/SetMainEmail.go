package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

// 设置主要的 Email

func SetMainEmail(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("设置Email失败", "设备异常"))
	}
	return c.JSON(result.Succeed.WithMsg("设置主要的Email"))
}
