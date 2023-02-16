package coinAI

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/router/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/reqCoinAI"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RemoveCoinAIParam struct {
	ServeID  string
	Password string
}

func Remove(c *fiber.Ctx) error {
	var json RemoveCoinAIParam
	mFiber.Parser(c, &json)

	if len(json.ServeID) < 3 {
		return c.JSON(result.Fail.WithMsg("缺少 ServeID"))
	}

	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: userID,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AItrade",
	}).Connect().Collection("CoinAINet")
	defer db.Close()

	findOpt := options.FindOne()
	findOpt.SetSort(map[string]int{
		"TimeUnix": -1,
	})

	findFK := bson.D{{
		Key:   "ServeID",
		Value: json.ServeID,
	}}
	var CoinServe dbType.AppEnv
	db.Table.FindOne(db.Ctx, findFK, findOpt).Decode(&CoinServe)
	if len(CoinServe.ServeID) < 3 {
		return c.JSON(result.Fail.WithMsg("未找到该服务"))
	}
	if CoinServe.UserID != userID {
		return c.JSON(result.Fail.WithMsg("该 CoinAI 不属于当前用户"))
	}

	// 在这里 ping 一下
	Origin := mStr.Join("http://", CoinServe.ServeID)
	_, err = reqCoinAI.NewRest(reqCoinAI.RestOpt{
		Origin: Origin,
		UserID: userID,
		Path:   "/ping",
		Method: "GET",
	})

	if err != nil {
		db.Table.DeleteOne(db.Ctx, findFK)
		return c.JSON(result.Succeed.WithMsg("已删除"))
	} else {
		return c.JSON(result.Fail.WithMsg("请先停止服务再进行删除！"))
	}
}
