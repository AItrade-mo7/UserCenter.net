package api

import (
	"fmt"

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

	Path := installShell.CoinFund(installShell.InstShellOpt{
		Port:   json.Port,
		UserID: json.UserID,
	})

	fmt.Println(Path)

	return c.JSON(result.Succeed.WithData(Path))
}
