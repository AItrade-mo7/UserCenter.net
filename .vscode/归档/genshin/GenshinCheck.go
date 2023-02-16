package account

import (
	"fmt"
	"strings"

	"UserCenter.net/server/genshin"
	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/router/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenshinCheckParam struct {
	Cookie string
}

func GenshinCheck(c *fiber.Ctx) error {
	var json GenshinCheckParam
	mFiber.Parser(c, &json)

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

	CookieStr := strings.TrimSpace(json.Cookie)

	resData, resErr := genshin.SignIn(CookieStr)

	if resErr != nil {
		return c.JSON(result.Fail.WithData(mStr.ToStr(resErr) + resData))
	}

	// 读取米游社 表
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AItrade",
	}).Connect().Collection("MiYouSheCookie")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("MiYouSheCookie,数据库连接错误 %+v", err)
		return c.JSON(result.Fail.WithData(resErr))
	}

	var dbRes dbType.MiYouSheCookieTable
	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.AccountData.UserID,
	}}
	db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)

	// 如果 Cookie 已存在 则直接返回
	if dbRes.MiYouSheCookie == CookieStr {
		return c.JSON(result.Succeed.WithData(string(resData) + "&&& 当前 Cookie 未改变"))
	}

	// 在这里更新到数据库里
	dbRes.Email = UserDB.AccountData.Email
	dbRes.UserID = UserDB.AccountData.UserID
	dbRes.CreateTime = mTime.GetUnixInt64()
	dbRes.CreateTimeStr = mTime.UnixFormat(dbRes.CreateTime)
	dbRes.MiYouSheCookie = CookieStr

	UK := bson.D{}
	mStruct.Traverse(dbRes, func(key string, val any) {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   key,
					Value: val,
				},
			},
		})
	})

	upOpt := options.Update()
	upOpt.SetUpsert(true)
	_, err = db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
	if err != nil {
		global.LogErr("GenshinCheck,数据更插失败", err)
		return c.JSON(result.ErrDB.WithData("数据更插失败"))
	}

	return c.JSON(result.Succeed.WithData(string(resData) + "&&& 当前 Cookie 已被添加到数据库定时队列。"))
}

func GetGenshinCookie(c *fiber.Ctx) error {
	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	// 读取米游社 表
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AItrade",
	}).Connect().Collection("MiYouSheCookie")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("MiYouSheCookie,数据库连接错误 %+v", err)
		return resErr
	}

	var dbRes dbType.MiYouSheCookieTable
	FK := bson.D{{
		Key:   "UserID",
		Value: UserID,
	}}
	db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)

	return c.JSON(result.Succeed.WithData(dbRes))
}
