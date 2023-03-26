package coinAI

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PublicList(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("获取失败", "设备异常"))
	}

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AIServe",
	}).Connect().Collection("CoinAI")
	defer db.Close()

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{{
		Key:   "IsPublic",
		Value: true,
	}}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}

	CoinAIList := []dbType.AppEnvType{}
	for cursor.Next(db.Ctx) {
		var CoinServe dbType.AppEnvType
		cursor.Decode(&CoinServe)
		CoinAIList = append(CoinAIList, CoinServe)
	}
	return c.JSON(result.Succeed.WithData(CoinAIList))
}
