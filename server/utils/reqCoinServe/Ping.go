package reqCoinServe

import (
	"fmt"

	"DataCenter.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mStr"
)

type AITradePingData struct {
	AppInfo struct {
		Name    string `bson:"name"`
		Version string `bson:"version"`
		Port    int    `bson:"Port"`
	} `bson:"AppInfo"`
	ContentType string   `bson:"ContentType"`
	FullPath    string   `bson:"FullPath"`
	Method      string   `bson:"Method"`
	ResParam    struct{} `bson:"ResParam"`
	Token       string   `bson:"Token"`
	UserAgent   string   `bson:"UserAgent"`
}

type AITradePingResult struct {
	Code int             `bson:"Code"`
	Data AITradePingData `bson:"Data"`
	Msg  string          `bson:"Msg"`
}

type PingOpt struct {
	ServerInfo  dbType.CoinServeTable
	AccountData dbType.AccountTable
}

func Ping(opt PingOpt) (resErr error) {
	resErr = nil

	pingBaseUrl := mStr.Join(
		"http://",
		opt.ServerInfo.CoinServeID,
	)

	_, err := NewRest(RestOpt{
		Origin: pingBaseUrl,
		UserID: opt.AccountData.UserID,
		Path:   "/CoinAI/ping",
		Method: "GET",
	})
	if err != nil {
		resErr = fmt.Errorf("服务验证失败")
		return
	}

	return
}
