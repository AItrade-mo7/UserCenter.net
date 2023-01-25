package account

import (
	"fmt"

	"DataCenter.net/server/genshin"
	"DataCenter.net/server/global"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

	resData, resErr := genshin.SignIn(json.Cookie)

	if resErr != nil {
		return c.JSON(result.Fail.WithData(resErr))
	}
	// 更新至 数据库
	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.UserID,
	}}
	UK := bson.D{}
	UK = append(UK, bson.E{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "MiYouSheCookie",
				Value: json.Cookie,
			},
		},
	})

	_, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	if err != nil {
		errStr := fmt.Errorf("数据库更新失败 %+v", err)
		global.LogErr(errStr)
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(errStr)))
	}
	UserDB.Update()

	return c.JSON(result.Succeed.WithData(string(resData) + "&&& 当前 Cookie 已被添加到数据库。"))
}
