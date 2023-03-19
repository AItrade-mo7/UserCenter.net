package coinAI

import (
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func PublicList(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("获取失败", "设备异常"))
	}

	// UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
	// 	Email: "config.SysEmail",
	// })
	// if err != nil {
	// 	return c.JSON(result.ErrDB.WithData(err))
	// }

	// db := mMongo.New(mMongo.Opt{
	// 	UserName: config.SysEnv.MongoUserName,
	// 	Password: config.SysEnv.MongoPassword,
	// 	Address:  config.SysEnv.MongoAddress,
	// 	DBName:   "AItrade",
	// }).Connect().Collection("CoinAINet")
	// defer db.Close()

	// findOpt := options.Find()
	// findOpt.SetAllowDiskUse(true)
	// findOpt.SetSort(map[string]int{
	// 	"TimeUnix": -1,
	// })

	// findFK := bson.D{{
	// 	Key:   "UserID",
	// 	Value: UserDB.AccountData.UserID,
	// }}

	// cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	// if err != nil {
	// 	return c.JSON(result.ErrDB.WithData(err))
	// }

	// CoinAIList := []dbType.AppEnv{}
	// for cursor.Next(db.Ctx) {
	// 	var CoinServe dbType.AppEnv
	// 	cursor.Decode(&CoinServe)
	// 	CoinAIList = append(CoinAIList, CoinServe)
	// }

	// return c.JSON(result.Succeed.WithData(CoinAIList))
	return c.JSON(result.Succeed.WithData([]dbType.AppEnvType{}))
}
