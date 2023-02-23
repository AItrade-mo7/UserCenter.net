package account

import (
	"fmt"
	"time"

	"UserCenter.net/server/global/apiType"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mRes"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type LoginParam struct {
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

func Login(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("登录失败", "设备异常"))
	}

	var json LoginParam
	mFiber.Parser(c, &json)

	// 验证邮箱和密码
	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		return c.JSON(result.ErrLogin.With("邮箱格式不正确", json.Email))
	}

	if len(json.Password) != 32 {
		return c.JSON(result.ErrLogin.With("密码格式不正确", "可能原因:密码没有加密传输！"))
	}

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		Email: json.Email,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()

	if len(UserDB.UserID) != 32 {
		UserDB.DB.Close()
		return c.JSON(result.ErrAccount.WithData("该邮箱尚未注册"))
	}

	IPInfoList := mVerify.GetIPS(c.IPs())
	var IPInfo mVerify.IPAddressType
	if len(IPInfoList) > 0 {
		IPInfo = IPInfoList[0]
	}
	if len(IPInfo.Hostname) > 0 {
		mJson.Println(IPInfo)
	}

	NewToken := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Message:   UserDB.Data.UserID,
		Issuer:    "AItrade.net",
		Subject:   "UserToken",
	}).Generate()
	LoginInfo := apiType.LoginSucceedType{
		Token: NewToken,
	}

	taskPush.SysEmail(taskPush.SysEmailOpt{
		To:             []string{json.Email},
		Subject:        "您已登录",
		Title:          "登录成功",
		Message:        "新的令牌已生成，登录设备信息：",
		Content:        c.Get("User-Agent"),
		Description:    "账号登录",
		EntrapmentCode: UserDB.Data.EntrapmentCode,
	})

	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	UserDB.DB.Close()

	// 需要在这里把 登录信息 存起来
	// db := mMongo.New(mMongo.Opt{
	// 	UserName: config.SysEnv.MongoUserName,
	// 	Password: config.SysEnv.MongoPassword,
	// 	Address:  config.SysEnv.MongoAddress,
	// 	DBName:   "AItrade",
	// }).Connect().Collection("Token")
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	db.Close()
	// 	resErr := fmt.Errorf("MiYouSheCookie,数据库连接错误 %+v", err)
	// 	return c.JSON(result.ErrDB.WithData(resErr))
	// }
	// var dbRes dbType.TokenTable
	// FK := bson.D{{
	// 	Key:   "UserID",
	// 	Value: loginSucceedData.UserID,
	// }}

	// db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)
	// dbRes.Token = loginSucceedData.Token
	// dbRes.UserID = loginSucceedData.UserID
	// dbRes.CreateTime = mTime.GetUnixInt64()

	// UK := bson.D{}
	// mStruct.Traverse(dbRes, func(key string, val any) {
	// 	UK = append(UK, bson.E{
	// 		Key: "$set",
	// 		Value: bson.D{
	// 			{
	// 				Key:   key,
	// 				Value: val,
	// 			},
	// 		},
	// 	})
	// })

	// upOpt := options.Update()
	// upOpt.SetUpsert(true)
	// _, err = db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
	// if err != nil {
	// 	global.LogErr("Login Token,数据更插失败", err)
	// 	return c.JSON(result.ErrDB.WithData("数据更插失败"))
	// }

	// return c.JSON(result.RightLogin.WithData(loginSucceedData))
	return c.JSON(result.RightLogin.WithData(LoginInfo))
}

type IPAddressType struct {
	ISP      string
	Hostname string
	Country  string
	Region   string
	City     string
}

func IPAddr(ips []string) (resData []IPAddressType, resErr error) {
	ipArr := struct {
		IP []string
	}{
		IP: ips,
	}
	res, err := taskPush.Request(taskPush.RequestOpt{
		Origin: config.SysEnv.MessageBaseUrl,
		Path:   "/api/public/GetAddressIP",
		Data:   mJson.ToJson(ipArr),
	})
	if err != nil {
		resErr = err
		return
	}

	var resObj mRes.ResType
	jsoniter.Unmarshal(res, &resObj)
	if resObj.Code < 0 {
		resErr = fmt.Errorf(resObj.Msg)
		return
	}

	fmt.Println(resObj.Data)

	return
}
