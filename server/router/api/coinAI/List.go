package coinAI

import (
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	// userID, err := middle.TokenAuth(c)
	// if err != nil {
	// 	return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
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
	// 	Value: userID,
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
	return c.JSON(result.Succeed.WithData("CoinAIList"))
}
