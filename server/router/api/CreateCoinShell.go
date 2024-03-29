package api

import (
	"UserCenter.net/server/utils/installShell"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
)

type InstallCoinAIParam struct {
	Port   string `bson:"Port"`
	UserID string `bson:"UserID"` // 用户 ID
}

func InstallCoinAI(c *fiber.Ctx) error {
	var json InstallCoinAIParam
	mFiber.Parser(c, &json)

	if !mVerify.IsPort(json.Port) {
		json.Port = "9856"
	}

	Path := installShell.CoinFund(installShell.InstShellOpt{
		Port:   json.Port,
		UserID: json.UserID,
	})

	return c.SendFile(Path)
}
