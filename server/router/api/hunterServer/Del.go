package hunterServer

import (
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/ReqHunterNet"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type DelParam struct {
	HunterServerID string `bson:"HunterServerID"`
	Password       string `bson:"Password"`
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
	ServerDB, err := LineHunterServer()
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	FK := bson.D{{
		Key:   "HunterServerID",
		Value: json.HunterServerID,
	}}

	var ServerData dbType.HunterServer

	ServerDB.Table.FindOne(ServerDB.Ctx, FK).Decode(&ServerData)

	if len(ServerData.HunterServerID) < 6 {
		return c.JSON(result.Fail.WithMsg("该服务不存在"))
	}

	err = ReqHunterNet.Ping(ReqHunterNet.PingOpt{
		ServerInfo:  ServerData,
		AccountData: UserInfo.AccountData,
	})

	if err == nil {
		return c.JSON(result.Fail.WithMsg("当前服务正在运行"))
	}

	FK = bson.D{{
		Key:   "HunterServerID",
		Value: json.HunterServerID,
	}}
	_, err = ServerDB.Table.DeleteOne(ServerDB.Ctx, FK)
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	ServerDB.Close()
	return c.JSON(result.Succeed.WithData("删除 Server"))
}
