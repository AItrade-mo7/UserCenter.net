package okxKey

import (
	"net"

	"DataCenter.net/server/global/apiType"
	"DataCenter.net/server/okxApi/testOkxApi"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
)

func AddOkxKey(c *fiber.Ctx) error {
	var json apiType.CreateOkxKeyParam
	mFiber.Parser(c, &json)

	// 检查 Name 是否合法
	isName := mVerify.IsNickName(json.Name)
	if !isName {
		return c.JSON(result.ErrAddOkxKey.With("Name不合法", "1-12位字母数字下划线"))
	}

	// 检查 IP 是否合法
	address := net.ParseIP(json.IP)
	if address == nil {
		return c.JSON(result.ErrAddOkxKey.WithMsg("ip地址不合法"))
	}

	switch {
	case len(json.ApiKey) < 30:
		return c.JSON(result.ErrAddOkxKey.WithMsg("请输入有效的 ApiKey"))
	case len(json.SecretKey) < 30:
		return c.JSON(result.ErrAddOkxKey.WithMsg("请输入有效的 SecretKey"))
	case len(json.Passphrase) < 4:
		return c.JSON(result.ErrAddOkxKey.WithMsg("请输入有效的 Passphrase"))
	}

	// 在这里检查 ApiKey
	_, err := testOkxApi.Start(testOkxApi.Opt{
		ApiKey:     json.ApiKey,
		Passphrase: json.Passphrase,
		SecretKey:  json.SecretKey,
	})
	if err != nil {
		return c.JSON(result.ErrAddOkxKey.WithMsg("ApiKey 检查失败"))
	}

	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	err = UserDB.CreateOkxKey(json)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithData(UserDB.AccountData.OkxKeyList))
}
