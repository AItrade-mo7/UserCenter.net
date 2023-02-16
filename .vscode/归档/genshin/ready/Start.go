package ready

import (
	"fmt"
	"time"

	"UserCenter.net/server/genshin"
	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/global/dbType"
	"UserCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	go mClock.New(mClock.OptType{
		Func: MiYouSheSign,
		Spec: "0 0 1,10,15,21 * * ? ", // 每天的 1,10,15,21
	})
}

func MiYouSheSign() {
	// 读取米游社 表
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AItrade",
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

		time.Sleep(time.Second * 2)
	}
}
