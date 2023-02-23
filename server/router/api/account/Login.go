package account

import (
	"time"

	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/dbUser"
	"UserCenter.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// 检测账号与密码
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

	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	// 生成Token
	NewToken := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Message:   UserDB.Data.UserID,
		Issuer:    "AItrade.net",
		Subject:   "UserToken",
	}).Generate()

	// 生成 登录信息
	DeviceInfo := mVerify.DeviceToUA(c.Get("User-Agent"))
	IPInfoList := mVerify.GetIPS(c.IPs())
	var IPInfo mVerify.IPAddressType
	if len(IPInfoList) > 0 {
		IPInfo = IPInfoList[0]
	}

	nowTime := mTime.GetTime()
	LoginInfo := dbType.LoginSucceedType{
		UserID:         UserDB.Data.UserID,
		Email:          UserDB.Data.Email,
		BrowserName:    DeviceInfo.BrowserName,
		OsName:         DeviceInfo.OsName,
		Hostname:       IPInfo.Hostname,
		ISP:            IPInfo.ISP,
		Operators:      IPInfo.Operators,
		CreateTimeUnix: nowTime.TimeUnix,
		CreateTimeStr:  nowTime.TimeStr,
		Token:          NewToken,
	}
	UserDB.DB.Close()

	// 生成邮件
	EmailCont := mStr.Join(
		"<br />",
		"时间: ", LoginInfo.CreateTimeStr, "<br />",
		"地区: ", LoginInfo.ISP, "<br />",
		"运营商: ", LoginInfo.Operators, "<br />",
		"系统: ", LoginInfo.OsName, "<br />",
		"设备: ", LoginInfo.BrowserName, "<br />",
		"IP: ", LoginInfo.Hostname, "<br />",
	)

	taskPush.SysEmail(taskPush.SysEmailOpt{
		To:             []string{json.Email},
		Subject:        "登录信息",
		Title:          "令牌已生成",
		Message:        "系统检测到如下登录信息:",
		Content:        EmailCont,
		Description:    "登录邮件",
		EntrapmentCode: UserDB.Data.EntrapmentCode,
	})

	// 登录记录存储
	dbLogin := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Account",
	}).Connect().Collection("LoginInfo")
	defer dbLogin.Close()
	_, err = dbLogin.Table.InsertOne(dbLogin.Ctx, LoginInfo)
	if err != nil {
		global.LogErr("account.Login, 登录信息存储失败", err)
		return c.JSON(result.ErrDB.WithData("Token存储失败"))
	}
	dbLogin.Close()

	// 验证 Token 的存储
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("VerifyToken")
	defer db.Close()

	FK := bson.D{{
		Key:   "UserID",
		Value: LoginInfo.UserID,
	}}

	var dbRes dbType.TokenTable
	db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)
	dbRes.UserID = LoginInfo.UserID
	dbRes.Token = LoginInfo.Token
	dbRes.CreateTime = mTime.GetUnixInt64()

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
		global.LogErr("account.Login, Token存储失败", err)
		return c.JSON(result.ErrDB.WithData("Token存储失败"))
	}

	return c.JSON(result.RightLogin.WithData(LoginInfo))
}
