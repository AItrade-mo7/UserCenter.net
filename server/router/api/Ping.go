package api

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type AppInfoType struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
}

func Ping(c *fiber.Ctx) error {
	json := mFiber.Parser(c)

	// 在这里请求数据
	ClientFileReqData, _ := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://raw.githubusercontent.com",
		Path:   "/mo7cc/trade.mo7.cc/main/package.json?tmp=" + mTime.GetUnix(),
	}).Get()

	CoinAIFileReqData, _ := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://raw.githubusercontent.com",
		Path:   "/AItrade-mo7/CoinAIPackage/main/package.json?tmp=" + mTime.GetUnix(),
	}).Get()

	var ClientInfo AppInfoType
	jsoniter.Unmarshal(ClientFileReqData, &ClientInfo)

	var CoinAIInfo AppInfoType
	jsoniter.Unmarshal(CoinAIFileReqData, &ClientInfo)

	var ApiInfo AppInfoType
	jsoniter.Unmarshal(mJson.ToJson(config.AppInfo), &ApiInfo)

	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json
	ReturnData["Method"] = c.Method()
	ReturnData["ApiInfo"] = ApiInfo
	ReturnData["ClientInfo"] = ClientInfo
	ReturnData["CoinAIInfo"] = CoinAIInfo

	ReturnData["UserAgent"] = c.Get("User-Agent")
	ReturnData["Path"] = c.OriginalURL()

	DeviceInfo := mVerify.DeviceToUA(c.Get("User-Agent"))
	ReturnData["BrowserName"] = DeviceInfo.BrowserName
	ReturnData["OsName"] = DeviceInfo.OsName

	ips := c.IPs()
	if len(ips) > 0 {
		ReturnData["IP"] = ips[0]
	}

	// 获取 token
	token := c.Get("Token")
	if len(token) > 0 {
		// Token 验证
		_, err := middle.TokenAuth(c)
		if err != nil {
			return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
		}
		ReturnData["Token"] = token
	}

	return c.JSON(result.Succeed.WithData(ReturnData))
}
