package coinAI

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/sysPublic/dbType"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func List(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("获取失败", "设备异常"))
	}

	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AIServe",
	}).Connect()
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer db.Close()
	db.Collection("CoinAI")

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{{
		Key:   "UserID",
		Value: userID,
	}}

	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}

	CoinAIList := []dbType.CoinAIType{}
	for cursor.Next(db.Ctx) {
		var CoinServe dbType.CoinAIType
		cursor.Decode(&CoinServe)
		CoinAIList = append(CoinAIList, CoinServe)
	}

	return c.JSON(result.Succeed.WithData(CoinAIList))
}
