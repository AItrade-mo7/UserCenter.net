package installShell

import (
	"bytes"
	"os"
	"text/template"

	"DataCenter.net/server/global/config"
	"DataCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

// 生成外部脚本
type InstShellOpt struct {
	Port   string
	UserID string
}

func CoinFund(opt InstShellOpt) (resData string, resErr error) {
	resErr = nil
	resData = ""

	savePath := mStr.Join(
		config.Dir.JsonData,
		mStr.ToStr(os.PathSeparator),
		"install",
	)

	name := mEncrypt.RandStr(1)
	fileName := mFile.GetName(mFile.GetNameOpt{
		FileName: mStr.Join(name, "-", opt.Port, ".sh"),
		SavePath: savePath,
	})
	filePath := mStr.Join(
		savePath,
		mStr.ToStr(os.PathSeparator),
		fileName,
	)

	// 生成文件
	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(tmpl.InstCoinServe))
	Tmpl.Execute(Body, tmpl.InstCoinServeParam{
		Port:   opt.Port,
		UserID: opt.UserID,
	})
	Cont := Body.String()

	// 写入文件
	mFile.Write(filePath, Cont)

	resData = filePath

	return
}
