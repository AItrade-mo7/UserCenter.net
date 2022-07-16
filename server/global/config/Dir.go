package config

import (
	"os"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

type DirType struct {
	Home      string // Home 根目录
	App       string // APP 根目录
	Log       string // 日志文件目录
	AITrade   string // AITrade 目录
	File      string // AITrade 目录
	FilRemote string // 静态文件服务器域名
}

var Dir DirType

type FileType struct {
	SysEnv      string // /root/sys_env.yaml
	LocalSysEnv string // ./sys_env.yaml
	AppEnv      string // ./app_env.yaml
}

var File FileType

func DirInit() {
	Dir.Home = mPath.HomePath()

	Dir.App, _ = os.Getwd()

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)

	Dir.File = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"ProdProject",
		mStr.ToStr(os.PathSeparator),
		"file.mo7.cc",
	)

	Dir.AITrade = mStr.Join(
		Dir.File,
		mStr.ToStr(os.PathSeparator),
		"AITrade",
	)

	Dir.FilRemote = "//file.mo7.cc"

	File.SysEnv = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)
	File.LocalSysEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)
	File.AppEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"app_env.yaml",
	)
}
