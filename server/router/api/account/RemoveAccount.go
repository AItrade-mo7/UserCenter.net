package account

import (
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/taskPush"
	"UserCenter.net/sysPublic/dbType"
	"UserCenter.net/sysPublic/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 设置主要的 Email
type RemoveAccountParam struct {
	EmailCode string
	Password  string
}

func RemoveAccount(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("设置Email失败", "设备异常"))
	}

	var json RemoveAccountParam
	mFiber.Parser(c, &json)

	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()
	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}
	// 验证码验证
	err = taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: UserDB.Data.Email,
		Code:  json.EmailCode,
	})
	if err != nil {
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}
	// 检查未解绑邮箱
	if len(UserDB.Data.UserEmail) > 1 {
		return c.JSON(result.Fail.WithMsg("存在未解绑的邮箱"))
	}

	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AIServe",
	}).Connect()
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}
	defer db.Close()
	db.Collection("CoinAI")

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	// 检查ApiKey
	findFK := bson.D{{
		Key:   "ApiKeyList.UserID",
		Value: UserID,
	}}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}
	CoinAIList := []dbType.AppEnvType{}
	for cursor.Next(db.Ctx) {
		var CoinServe dbType.AppEnvType
		cursor.Decode(&CoinServe)
		CoinAIList = append(CoinAIList, CoinServe)
	}

	if len(CoinAIList) > 0 {
		return c.JSON(result.Fail.WithMsg("存在未删除的ApiKey"))
	}

	// 检查 卫星服务
	findFK2 := bson.D{{
		Key:   "UserID",
		Value: UserID,
	}}
	cursor2, err := db.Table.Find(db.Ctx, findFK2, findOpt)
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}
	CoinAIList2 := []dbType.AppEnvType{}
	for cursor2.Next(db.Ctx) {
		var CoinServe dbType.AppEnvType
		cursor2.Decode(&CoinServe)
		CoinAIList2 = append(CoinAIList2, CoinServe)
	}

	if len(CoinAIList2) > 0 {
		return c.JSON(result.Fail.WithMsg("存在未删除的卫星服务!"))
	}

	// 删除账户
	dbAccount, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Account",
	}).Connect()
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}
	defer dbAccount.Close()
	dbAccount.Collection("User")

	FK := bson.D{{
		Key:   "UserID",
		Value: UserID,
	}}
	_, err = dbAccount.Table.DeleteOne(dbAccount.Ctx, FK)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	taskPush.DelEmailCode(UserDB.Data.Email)

	taskPush.SysEmail(taskPush.SysEmailOpt{
		To:             UserDB.Data.UserEmail,
		Subject:        "您的账户已被删除",
		Title:          "刚刚您的账户被注销了",
		Message:        "你的账户所有数据均被删除和清空！",
		Content:        "感谢您对系统的支持，若有任何修改意见可发送邮件至 meichangliang@outlook.com ",
		Description:    "注销账户通知",
		EntrapmentCode: UserDB.Data.EntrapmentCode,
	})

	return c.JSON(result.Succeed.WithData("账户已删除"))
}
