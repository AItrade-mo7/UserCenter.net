package api

import (
	"fmt"

	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mRes"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type SendEmailCodeParam struct {
	Email          string
	Action         string
	EntrapmentCode string // 非必填
}

func SendEmailCode(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("发送失败", "设备异常"))
	}

	var json SendEmailCodeParam
	mFiber.Parser(c, &json)

	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		emailErr := fmt.Errorf("邮箱格式不正确 %+v", json.Email)
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	if len(json.Action) < 1 {
		emailErr := fmt.Errorf("Action不能为空")
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	// 优先去数据库寻找防钓鱼码
	Token := c.Get("Token") // 如果登录了，则使用登录的验证码
	if len(Token) > 0 {
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
		defer UserDB.DB.Close()
		// 如果存在防钓鱼码则优先使用 数据库的 否则使用传入的
		if len(UserDB.Data.EntrapmentCode) > 0 {
			json.EntrapmentCode = UserDB.Data.EntrapmentCode
		}
	}

	if len([]rune(json.EntrapmentCode)) > 24 {
		emailErr := fmt.Errorf("防钓鱼码不能大于24位")
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	resData, err := taskPush.Request(taskPush.RequestOpt{
		Origin: config.SysEnv.MessageBaseUrl,
		Path:   "/api/await/SendEmailCode",
		Data:   mJson.ToJson(json),
	})
	if err != nil {
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	var resObj mRes.ResType
	jsoniter.Unmarshal(resData, &resObj)

	if resObj.Code < 0 {
		return c.JSON(result.ErrEmailCode.WithMsg(resObj.Msg))
	}

	return c.JSON(result.Succeed.WithMsg("验证码已发送"))
}
