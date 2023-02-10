package msg

import (
	"fmt"

	"github.com/EasyGolang/goTools/mJson"
)

// var BaseUrl = "http://msg.mo7.cc"
var BaseUrl = "http://127.0.0.1:8900"

func SendMsg() {
	resData, resErr := NewReq(ReqOpt{
		Origin: BaseUrl,
		Path:   "/api/public/InsertTaskQueue",
		Method: "POST",
		Data: map[string]any{
			"jsonrpc": "2.0",
			"id":      "5",
		},
	})

	fmt.Println(mJson.JsonFormat(resData), resErr)
}
