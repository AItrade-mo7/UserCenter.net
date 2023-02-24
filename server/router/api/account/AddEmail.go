package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

// 新增 Email

func AddEmail(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("新增Email失败", "设备异常"))
	}
	return c.JSON(result.Succeed.WithMsg("新增Email"))
}
