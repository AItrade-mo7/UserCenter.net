package taskPush

import (
	"UserCenter.net/server/global/config"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
)

type ReqOpt struct {
	Origin string
	Path   string
	Method string
	Data   map[string]any
}

func Req(opt ReqOpt) (resData []byte, resErr error) {
	UserAgent := "AItrade.net"
	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: opt.Origin,
		Path:   opt.Path,
		Data:   mJson.ToJson(opt.Data),
		Header: map[string]string{
			"Auth-Encrypt": config.ClientEncrypt(opt.Path + UserAgent),
			"User-Agent":   UserAgent,
		},
	})

	if opt.Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
