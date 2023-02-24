package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

// 获取EmailList

func GetEmailList(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("获取EmailList失败", "设备异常"))
	}
	return c.JSON(result.Succeed.WithMsg("获取EmailList"))
}
