package hunterServer

import (
	"fmt"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/apiType"
	"DataCenter.net/server/global/config"
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/router/middle"
	"DataCenter.net/server/router/result"
	"DataCenter.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddServer(c *fiber.Ctx) error {
	var json apiType.CreateHunterServerParam
	mFiber.Parser(c, &json)

	// 检查端口
	if mCount.Le(json.Port, "79") < 0 || mCount.Le(json.Port, "60000") > 0 {
		return c.JSON(result.ErrAddHunterServer.WithMsg("端口号不合法"))
	}

	if len(json.OkxKeyID) < 10 {
		return c.JSON(result.ErrAddHunterServer.WithMsg("请选择有效的秘钥"))
	}

	if len(json.Note) < 3 {
		return c.JSON(result.ErrAddHunterServer.WithMsg("备注不能少于3个字"))
	}

	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		UserDB.DB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	// 密码验证
	err = UserDB.CheckPassword(json.Password)
	if err != nil {
		return c.JSON(result.ErrLogin.WithMsg(err))
	}

	if len(UserDB.AccountData.OkxKeyList) < 1 {
		return c.JSON(result.ErrAddHunterServer.WithMsg("账户中没有秘钥"))
	}

	// 查找秘钥
	OkxKey := dbType.OkxKey{}
	for _, val := range UserDB.AccountData.OkxKeyList {
		if val.OkxKeyID == json.OkxKeyID {
			OkxKey = val
		}
	}

	if len(OkxKey.OkxKeyID) < 10 {
		return c.JSON(result.ErrAddHunterServer.WithMsg("该秘钥不存在"))
	}
	UserDB.DB.Close()

	// 构建 HunterServer 数据结构
	HunterServerData := dbType.HunterServer{}
	HunterServerData.Host = OkxKey.IP
	HunterServerData.Port = json.Port
	HunterServerData.OkxKeyID = OkxKey.OkxKeyID
	HunterServerData.UserID = UserDB.UserID
	HunterServerData.Note = json.Note
	HunterServerData.CreateTime = mTime.GetUnixInt64()
	HunterServerData.HunterServerID = mStr.Join(
		OkxKey.IP, ":", json.Port,
	)

	// 连接 HunterServer 表
	ServerDB, err := LineHunterServer()
	if err != nil {
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}

	// 检查服务是否存在  --  HunterServerID
	FK := bson.D{{
		Key:   "HunterServerID",
		Value: HunterServerData.HunterServerID,
	}}
	var ServerDB_data dbType.HunterServer
	ServerDB.Table.FindOne(ServerDB.Ctx, FK).Decode(&ServerDB_data)
	if len(ServerDB_data.HunterServerID) > 2 {
		return c.JSON(result.ErrAddHunterServer.With("该服务已存在", "建议更换端口"))
	}

	// 检查服务是否存在  --  OkxKeyID
	FK = bson.D{{
		Key:   "OkxKeyID",
		Value: HunterServerData.OkxKeyID,
	}}
	ServerDB_data = dbType.HunterServer{}
	ServerDB.Table.FindOne(ServerDB.Ctx, FK).Decode(&ServerDB_data)
	if len(ServerDB_data.HunterServerID) > 2 {
		return c.JSON(result.ErrAddHunterServer.With("秘钥已被使用", "建议更换其它密钥"))
	}

	// 插入数据
	_, err = ServerDB.Table.InsertOne(ServerDB.Ctx, HunterServerData)
	if err != nil {
		errStr := fmt.Errorf("注册,插入数据失败 %+v", err)
		global.LogErr(errStr)
		ServerDB.Close()
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithData(HunterServerData))
}

func LineHunterServer() (resDB *mMongo.DB, resErr error) {
	resErr = nil
	resDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Hunter",
	}).Connect().Collection("HunterServer")

	err := resDB.Ping()
	if err != nil {
		resDB.Close()
		errStr := fmt.Errorf("用户数据读取失败,数据库连接错误 %+v", err)
		global.LogErr(errStr)
		resErr = errStr
		return
	}

	return
}
