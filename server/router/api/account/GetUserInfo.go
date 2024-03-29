package account

import (
	"UserCenter.net/server/global/apiType"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/sysPublic/dbUser"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func GetUserInfo(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("获取用户信息失败", "设备异常"))
	}

	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: userID,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()

	if len(UserDB.UserID) < 32 {
		return c.JSON(result.ErrToken.WithData("该用户不存在"))
	}

	var userinfoData apiType.UserInfo
	jsonStr := mJson.ToJson(UserDB.Data)
	jsoniter.Unmarshal(jsonStr, &userinfoData)

	return c.JSON(result.Succeed.WithData(userinfoData))
}
