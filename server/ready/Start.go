package ready

import (
	"fmt"
	"time"

	"DataCenter.net/server/genshin"
	"DataCenter.net/server/global"
	"DataCenter.net/server/global/config"
	"DataCenter.net/server/global/dbType"
	"DataCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	MiYouSheSign()
	go mClock.New(mClock.OptType{
		Func: MiYouSheSign,
		Spec: "0 0 1,6,11,16,20,23 * * ? ", // 每天的 1  10  16  23
	})
}

func MiYouSheSign() {
	// 读取米游社 表
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MiYouSheCookie")

	defer db.Close()
	err := db.Ping()
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("MiYouSheSign,数据库连接错误 %+v", err)
		global.LogErr(resErr)
		return
	}

	FK := bson.D{}
	db.Table.Find(db.Ctx, FK)
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	cursor, err := db.Table.Find(db.Ctx, FK, findOpt)
	if err != nil {
		global.LogErr("MiYouSheSign,数据库连接错误2 %+v", err)
		return
	}

	MiYouSheList := []dbType.MiYouSheCookieTable{}
	for cursor.Next(db.Ctx) {
		var MiYouShe dbType.MiYouSheCookieTable
		cursor.Decode(&MiYouShe)
		MiYouSheList = append(MiYouSheList, MiYouShe)
	}
	db.Close()

	for _, val := range MiYouSheList {

		resData, resErr := genshin.SignIn(val.MiYouSheCookie)

		if resErr != nil {
			global.Email(global.EmailOpt{
				To: []string{
					val.Email,
				},
				Subject:  "米游社自动签到未知错误",
				Template: tmpl.SysEmail,
				SendData: tmpl.SysParam{
					Message:      resData,
					SysTime:      mTime.UnixFormat(mTime.GetUnixInt64()),
					SecurityCode: "米游社自动签到",
				},
			})
		}
		global.Log.Println("执行一次签到", val.Email, resData)

		time.Sleep(time.Second * 2) // 2 秒钟 一次
	}
}
