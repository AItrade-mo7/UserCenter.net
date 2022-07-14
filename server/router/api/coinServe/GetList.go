package coinServe

import (
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetList(c *fiber.Ctx) error {
	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	ServerDB, err := LineCoinServerDB()
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	FK := bson.D{{
		Key:   "UserID",
		Value: UserID,
	}}

	cur, err := ServerDB.Table.Find(ServerDB.Ctx, FK)
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	var List []dbType.CoinServeTable
	for cur.Next(ServerDB.Ctx) {
		var result dbType.CoinServeTable
		cur.Decode(&result)
		List = append(List, result)
	}

	return c.JSON(result.Succeed.WithData(List))
}
