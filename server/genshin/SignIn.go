package genshin

import (
	"fmt"
	"strings"

	"DataCenter.net/server/global/config"
	"DataCenter.net/server/tmpl"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mShell"
	"github.com/EasyGolang/goTools/mStr"
)

/*

页面：

https://www.miyoushe.com/ys/accountCenter/postList?id=158179235

接口：

getUserFullInfo

Cookie:

	genshin.SignIn("_MHYUUID=2bd94c1b-ffee-4cdd-87e8-ff0de88d2ba1; DEVICEFP_SEED_ID=ccaa7f9909485911; DEVICEFP_SEED_TIME=1673447944577; _ga=GA1.1.983084773.1673447946; DEVICEFP=38d7ecfed3c4c; acw_tc=2f6fc10a16746628722334567e28ecb1eba7405bf267d0aebcdcd0ce6f25c8; LOGIN_PLATFORM_SWITCH_STATUS={%22bll8iq97cem8%22:{%22password_reset_entry%22:true%2C%22qr_login%22:false%2C%22sms_login_tab%22:true%2C%22pwd_login_tab%22:true}}; cookie_token_v2=v2_izeQ0yzuoQ5rCX2GBe8Gh14nAynFsuer6etqDFrGm0MXRC3QjOBXvmOjlJTwhg3g86u3R43XQ8roYm5fJ0nive_k_-tlx8s2jz8Fhby-j5PO6kZQ6nfB; account_mid_v2=0ccxa8x924_mhy; account_id_v2=158179235; ltoken_v2=v2_a-A1R9t30MPXIxgVR69SDrmOTgm1031rwb03T63OnKH0LhnU3l20slugzp1vZfBHVpbQHmpASRc2ICw4B3NV3pKRkgPa_yu_; ltmid_v2=0ccxa8x924_mhy; ltuid_v2=158179235; _ga_KS4J8TXSHQ=GS1.1.1674662871.3.1.1674664618.0.0.0")


*/

func SignIn(cookie string) (resData string, resErr error) {
	PyStr := tmpl.SignInPy

	CookieStr := strings.TrimSpace(cookie)

	if len(CookieStr) < 12 {
		resErr = fmt.Errorf("cookie 长度不足")
		return
	}

	PyThonPath := config.Dir.JsonData + "/SignIn.py"

	mFile.Write(PyThonPath, PyStr)

	TempConfig := map[string]string{
		"PyThonPath": PyThonPath,
		"Cookie":     CookieStr,
	}
	ShellCont := `
python3 ${PyThonPath} "${Cookie}"
`

	ShellCont = mStr.Temp(ShellCont, TempConfig)

	res, resErr := mShell.Run(ShellCont)
	if resErr != nil {
		return
	}

	resData = string(res)

	find := strings.Contains(resData, "旅行者")
	if find {
		resErr = nil
	} else {
		resErr = fmt.Errorf("米游社返回错误")
	}

	return
}
