package account

import (
	"fmt"
	"time"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/apiType"
	"DataCenter.net/server/global/config"
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/tmpl"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Login(c *fiber.Ctx) error {
	var json struct {
		Email    string `bson:"Email"`
		Password string `bson:"Password"`
	}
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

	if len(UserDB.UserID) != 32 {
		UserDB.DB.Close()
		return c.JSON(result.ErrAccount.WithData("该邮箱尚未注册"))
	}

	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	var loginSucceedData apiType.LoginSucceedType
	jsonStr := mJson.ToJson(UserDB.AccountData)
	jsoniter.Unmarshal(jsonStr, &loginSucceedData)

	loginSucceedData.Token = mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Message:   UserDB.AccountData.UserID,
		Issuer:    "AITrade.net",
		Subject:   "UserToken",
	}).Generate()

	go global.Email(global.EmailOpt{
		To: []string{
			loginSucceedData.Email,
		},
		Subject:  "登录提醒",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message:      "您刚刚执行了登录操作，登录设备: " + c.Get("User-Agent"),
			SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
			SecurityCode: UserDB.AccountData.SecurityCode,
		},
	}).Send()

	UserDB.DB.Close()

	// 需要在这里把 token 存起来
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("Token")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("MiYouSheCookie,数据库连接错误 %+v", err)
		return c.JSON(result.ErrDB.WithData(resErr))
	}
	var dbRes dbType.TokenTable
	FK := bson.D{{
		Key:   "UserID",
		Value: loginSucceedData.UserID,
	}}

	db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)
	dbRes.Token = loginSucceedData.Token
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
		global.LogErr("Login Token,数据更插失败", err)
		return c.JSON(result.ErrDB.WithData("数据更插失败"))
	}

	return c.JSON(result.RightLogin.WithData(loginSucceedData))
}
