package ReqHunterNet

import (
	"fmt"

	"DataCenter.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mStr"
)

type HunterPingData struct {
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

type HunterPingResult struct {
	Code int            `json:"Code"`
	Data HunterPingData `json:"Data"`
	Msg  string         `json:"Msg"`
}

type PingOpt struct {
	ServerInfo  dbType.HunterServer
	AccountData dbType.AccountTable
}

func Ping(opt PingOpt) (resErr error) {
	resErr = nil

	pingBaseUrl := mStr.Join(
		"http://",
		opt.ServerInfo.HunterServerID,
	)

	reqData := NewRest(NewRestOpt{
		Origin: pingBaseUrl,
		UserID: opt.AccountData.UserID,
		Path:   "/hunter_net/ping",
		Method: "GET",
	})

	if len(reqData) < 5 {
		resErr = fmt.Errorf("服务验证失败")
		return
	}

	return
}
