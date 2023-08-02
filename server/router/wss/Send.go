package wss

import (
	"UserCenter.net/server/router/result"
	"github.com/EasyGolang/goTools/mRes"
	"github.com/EasyGolang/goTools/mTime"
)

func Send() mRes.ResType {
	data := map[string]any{
		"DataSource": "UserCenter.net",
		"SysTime":    mTime.GetUnixInt64(),
	}

	return result.Succeed.WithData(data)
}
