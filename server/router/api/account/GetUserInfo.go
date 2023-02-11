package account

import (
	"UserCenter.net/server/global/apiType"
	"UserCenter.net/server/router/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func GetUserInfo(c *fiber.Ctx) error {
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

	if len(UserDB.UserID) != 32 {
		UserDB.DB.Close()
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	var userinfoData apiType.UserInfo
	jsonStr := mJson.ToJson(UserDB.AccountData)
	jsoniter.Unmarshal(jsonStr, &userinfoData)

	UserDB.DB.Close()
	return c.JSON(result.Succeed.WithData(userinfoData))
}
