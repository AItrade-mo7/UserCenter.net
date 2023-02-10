package msg

import (
	"fmt"

	"github.com/EasyGolang/goTools/mJson"
)

var BaseUrl = "http://msg.mo7.cc"

func SendMsg() {
	resData, resErr := NewReq(ReqOpt{
		Origin: BaseUrl,
		Path:   "/",
		Method: "POST",
		Data: map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      "5",
		},
	})

	fmt.Println(mJson.JsonFormat(resData), resErr)
}
