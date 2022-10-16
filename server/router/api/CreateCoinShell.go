package api

import (
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/installShell"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
)

type InstallCoinShellParam struct {
	Port   string `bson:"Port"`
	UserID string `bson:"UserID"` // 用户 ID
}

func InstallCoinShell(c *fiber.Ctx) error {
	var json InstallCoinShellParam
	mFiber.Parser(c, &json)

	if !mVerify.IsPort(json.Port) {
		json.Port = "9856"
	}

	Url, err := installShell.CoinFund(installShell.InstShellOpt{
		Port:   json.Port,
		UserID: json.UserID,
	})
	if err != nil {
		return c.JSON(result.Succeed.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithData(Url))
}
