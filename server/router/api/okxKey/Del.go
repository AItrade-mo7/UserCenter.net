package okxKey

import (
	"fmt"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/api/coinServe"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type DelParam struct {
	OkxKeyID string `bson:"OkxKeyID"`
	Password string `bson:"Password"`
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

	// 查询当前 Key 是否正在被使用
	ServerDB, err := coinServe.LineCoinServerDB()
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	FK := bson.D{{
		Key:   "OkxKeyID",
		Value: json.OkxKeyID,
	}}

	var ServerData dbType.CoinServeTable

	ServerDB.Table.FindOne(ServerDB.Ctx, FK).Decode(&ServerData)

	if len(ServerData.CoinServeID) > 6 {
		return c.JSON(result.Fail.WithMsg("有服务正在使用此Key"))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	List := UserDB.AccountData.OkxKeyList
	NewList := []dbType.OkxKeyTable{}
	for _, val := range List {
		if val.OkxKeyID != json.OkxKeyID {
			NewList = append(NewList, val)
		}
	}

	FK = bson.D{{
		Key:   "UserID",
		Value: UserDB.AccountData.UserID,
	}}

	UK := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "OkxKeyList",
				Value: NewList,
			},
		},
	}}

	_, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	if err != nil {
		errStr := fmt.Errorf("数据库更新失败 %+v", err)
		global.LogErr(errStr)
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(errStr)))
	}

	UserDB.Update()

	return c.JSON(result.Succeed.WithData(UserDB.AccountData.OkxKeyList))
}
