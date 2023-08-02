package installShell

import (
	"bytes"
	"os"
	"text/template"

	"UserCenter.net/server/global/config"
	"UserCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

// 生成外部脚本
type InstShellOpt struct {
	Port   string
	UserID string
}

func CoinFund(opt InstShellOpt) (resData string) {
	savePath := mStr.Join(
		config.Dir.JsonData,
		mStr.ToStr(os.PathSeparator),
		"install",
	)
	isJsonDataPath := mPath.Exists(savePath)
	if !isJsonDataPath {
		os.MkdirAll(savePath, 0o777)
	}

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

	return filePath
}
