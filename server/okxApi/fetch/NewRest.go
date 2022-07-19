package fetch

import (
	"fmt"
	"strings"

	"DataCenter.net/server/okxApi"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mUrl"
)

type AuthInfoType struct {
	APIKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

type NewOkxFetchOpt struct {
	ApiKey     string
	Passphrase string
	SecretKey  string
	Path       string
	Method     string
	Data       map[string]any
}

func NewRest(opt NewOkxFetchOpt) (resData []byte, resErr error) {
	Timestamp := mTime.IsoTime(true)
	ApiKey := opt.ApiKey
	SecretKey := opt.SecretKey
	Passphrase := opt.Passphrase
	Body := mJson.ToJson(opt.Data)

	signStr := mStr.Join(
		Timestamp,
		strings.ToUpper(opt.Method),
		opt.Path,
		string(Body),
	)

	if opt.Method == "GET" {
		Body = []byte("")

		urlO := mUrl.InitUrl(opt.Path)
		for key, val := range opt.Data {
			v := fmt.Sprintf("%+v", val)
			urlO.AddParam(key, v)
		}
		signPath := urlO.String()

		signStr = mStr.Join(
			Timestamp,
			strings.ToUpper(opt.Method),
			signPath,
			string(Body),
		)
	}

	Sign := mEncrypt.Sha256(signStr, SecretKey)

	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: okxApi.BaseUrl.Rest,
		Path:   opt.Path,
		Data:   opt.Data,
		Header: map[string]string{
			"OK-ACCESS-KEY":        ApiKey,
			"OK-ACCESS-SIGN":       Sign,
			"OK-ACCESS-TIMESTAMP":  Timestamp,
			"OK-ACCESS-PASSPHRASE": Passphrase,
		},
	})

	if opt.Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
