package coinAI

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

type RemoveCoinAIParam struct {
	ServeID   string
	Password  string
	EmailCode string
}

func Remove(c *fiber.Ctx) error {
	var json RemoveCoinAIParam
	mFiber.Parser(c, &json)
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.Fail.With("删除失败", "设备异常"))
	}

	if len(json.ServeID) < 1 {
		return c.JSON(result.Fail.WithMsg("缺少 ServeID"))
	}
	if len(json.Password) < 16 {
		return c.JSON(result.Fail.With("密码格式不正确", "可能原因:没有加密传输！"))
	}
	if len(json.EmailCode) != 32 {
		return c.JSON(result.Fail.With("验证码格式不正确", "可能原因:没有加密传输！"))
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
	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}
	// 验证验证码
	err = taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
		Email: UserDB.Data.Email,
		Code:  json.EmailCode,
	})
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}

	// 检测服务是否在使用
	Origin := mStr.Join("http://", json.ServeID)
	_, err = taskPush.Request(taskPush.RequestOpt{
		Origin: Origin,
		Path:   "/ping",
	})
	if err == nil {
		return c.JSON(result.Succeed.WithMsg("服务正在运行，请先手动关闭服务！"))
	}

	// 开始删除数据
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

	findOpt := options.FindOne()
	findOpt.SetSort(map[string]int{
		"TimeUnix": -1,
	})

	findFK := bson.D{{
		Key:   "ServeID",
		Value: json.ServeID,
	}}
	var CoinServe dbType.CoinAIType
	db.Table.FindOne(db.Ctx, findFK, findOpt).Decode(&CoinServe)
	if len(CoinServe.ServeID) < 3 {
		return c.JSON(result.Fail.WithMsg("未找到该服务"))
	}
	if CoinServe.UserID != userID {
		return c.JSON(result.Fail.WithMsg("该 CoinAI 不属于当前用户"))
	}
	db.Table.DeleteOne(db.Ctx, findFK)

	taskPush.DelEmailCode(UserDB.Data.Email)

	return c.JSON(result.Succeed.WithMsg("已删除"))
}
