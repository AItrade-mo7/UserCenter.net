package reqCoinAI

import (
	"time"

	"UserCenter.net/server/global/config"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
)

type RestOpt struct {
	Origin string
	UserID string
	Path   string
	Method string
	Data   map[string]any
}

func NewRest(opt RestOpt) (resData []byte, resErr error) {
	Token := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,              // key
		ExpiresAt: time.Now().Add(time.Hour / 2), // 过期时间 半小时
		Message:   opt.UserID,
		Issuer:    "AItrade.net",
		Subject:   "UserToken",
	}).Generate()

	UserAgent := "AItrade.net"

	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: opt.Origin,
		Path:   opt.Path,
		Data:   mJson.ToJson(opt.Data),
		Header: map[string]string{
			"Auth-Encrypt": config.ClientEncrypt(opt.Path + UserAgent),
			"Token":        Token,
			"User-Agent":   UserAgent,
		},
	})

	if opt.Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
