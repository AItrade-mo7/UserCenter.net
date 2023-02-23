package account

import (
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("获取用户信息失败", "设备异常"))
	}

	// userID, err := middle.TokenAuth(c)
	// if err != nil {
	// 	return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	// }

	// UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
	// 	UserID: userID,
	// })
	// if err != nil {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	// }

	// if len(UserDB.UserID) != 32 {
	// 	UserDB.DB.Close()
	// 	return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	// }

	// var userinfoData apiType.UserInfo
	// jsonStr := mJson.ToJson(UserDB.AccountData)
	// jsoniter.Unmarshal(jsonStr, &userinfoData)

	// UserDB.DB.Close()
	// return c.JSON(result.Succeed.WithData(userinfoData))
	return c.JSON(result.Succeed.WithData("userinfoData"))
}
