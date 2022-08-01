package coinServe

import (
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"DataCenter.net/server/utils/reqCoinServe"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type DelParam struct {
	CoinServeID string `bson:"CoinServeID"`
	Password    string `bson:"Password"`
}

func Del(c *fiber.Ctx) error {
	var json DelParam
	mFiber.Parser(c, &json)

	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	UserInfo, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	err = UserInfo.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(mStr.ToStr(err)))
	}

	// 在这里验证 当前服务的运行状态
	ServerDB, err := LineCoinServerDB()
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	FK := bson.D{{
		Key:   "CoinServeID",
		Value: json.CoinServeID,
	}}

	var ServerData dbType.CoinServeTable

	ServerDB.Table.FindOne(ServerDB.Ctx, FK).Decode(&ServerData)

	if len(ServerData.CoinServeID) < 6 {
		return c.JSON(result.Fail.WithMsg("该服务不存在"))
	}

	err = reqCoinServe.Ping(reqCoinServe.PingOpt{
		ServerInfo:  ServerData,
		AccountData: UserInfo.AccountData,
	})

	if err == nil {
		return c.JSON(result.Fail.WithMsg("当前服务正在运行"))
	}

	FK = bson.D{{
		Key:   "CoinServeID",
		Value: json.CoinServeID,
	}}
	_, err = ServerDB.Table.DeleteOne(ServerDB.Ctx, FK)
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	ServerDB.Close()
	return c.JSON(result.Succeed.WithData("删除 Server"))
}
