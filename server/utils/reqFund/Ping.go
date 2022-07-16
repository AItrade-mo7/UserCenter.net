package reqFund

import (
	"fmt"

	"DataCenter.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mStr"
)

type AITradePingData struct {
	AppInfo struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Port    int    `json:"Port"`
	} `json:"AppInfo"`
	ContentType string   `json:"ContentType"`
	FullPath    string   `json:"FullPath"`
	Method      string   `json:"Method"`
	ResParam    struct{} `json:"ResParam"`
	Token       string   `json:"Token"`
	UserAgent   string   `json:"UserAgent"`
}

type AITradePingResult struct {
	Code int             `json:"Code"`
	Data AITradePingData `json:"Data"`
	Msg  string          `json:"Msg"`
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

	reqData := NewRest(RestOpt{
		Origin: pingBaseUrl,
		UserID: opt.AccountData.UserID,
		Path:   "/AITrade_net/ping",
		Method: "GET",
	})

	if len(reqData) < 5 {
		resErr = fmt.Errorf("服务验证失败")
		return
	}

	return
}
