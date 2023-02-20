package middle

import (
	"path"
	"strings"

	"UserCenter.net/server/global/config"
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func Public(c *fiber.Ctx) error {
	// 添加访问头
	c.Set("Data-Path", config.SysName)

	findWss := strings.Contains(c.Path(), "/wss")
	if findWss {
		return c.Next()
	}

	// 链接无后缀名则需要验证
	filenameWithSuffix := path.Base(c.Path())
	fileSuffix := path.Ext(filenameWithSuffix)
	if len(fileSuffix) < 2 {
		// 授权验证
		err := EncryptAuth(c)
		if err != nil {
			return c.JSON(result.ErrAuth.WithData(mStr.ToStr(err)))
		}
	}

	return c.Next()
}
