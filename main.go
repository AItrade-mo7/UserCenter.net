package main

import (
	_ "embed"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/config"
	"DataCenter.net/server/router"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.ApiInfo)
	// 初始化系统参数
	global.Start()

	// 启动 http 监听服务
	router.Start()

	// resData, resErr := genshin.SignIn("_MHYUUID=2bd94c1b-ffee-4cdd-87e8-ff0de88d2ba1; DEVICEFP_SEED_ID=ccaa7f9909485911; DEVICEFP_SEED_TIME=1673447944577; _MHYUUID=2bd94c1b-ffee-4cdd-87e8-ff0de88d2ba1; _ga=GA1.1.983084773.1673447946; DEVICEFP=38d7ecfed3c4c; LOGIN_PLATFORM_SWITCH_STATUS={%22bll8iq97cem8%22:{%22password_reset_entry%22:true%2C%22qr_login%22:false%2C%22sms_login_tab%22:true%2C%22pwd_login_tab%22:true}}; account_mid_v2=0ccxa8x924_mhy; account_id_v2=158179235; ltmid_v2=0ccxa8x924_mhy; ltuid_v2=158179235; _ga_KS4J8TXSHQ=GS1.1.1674674106.5.1.1674674116.0.0.0")

	// fmt.Println((resData), resErr)
}
