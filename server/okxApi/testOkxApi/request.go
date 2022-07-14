package testOkxApi

import (
	"fmt"

	"DataCenter.net/server/global"
	"DataCenter.net/server/okxApi/fetch"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type ReqType struct {
	Code string `json:"code"`
	Data []struct {
		AdjEq   string `json:"adjEq"`
		Details []struct {
			AvailBal      string `json:"availBal"`
			AvailEq       string `json:"availEq"`
			CashBal       string `json:"cashBal"`
			Ccy           string `json:"ccy"`
			CrossLiab     string `json:"crossLiab"`
			DisEq         string `json:"disEq"`
			Eq            string `json:"eq"`
			EqUsd         string `json:"eqUsd"`
			FrozenBal     string `json:"frozenBal"`
			Interest      string `json:"interest"`
			IsoEq         string `json:"isoEq"`
			IsoLiab       string `json:"isoLiab"`
			IsoUpl        string `json:"isoUpl"`
			Liab          string `json:"liab"`
			MaxLoan       string `json:"maxLoan"`
			MgnRatio      string `json:"mgnRatio"`
			NotionalLever string `json:"notionalLever"`
			OrdFrozen     string `json:"ordFrozen"`
			StgyEq        string `json:"stgyEq"`
			Twap          string `json:"twap"`
			UTime         string `json:"uTime"`
			Upl           string `json:"upl"`
			UplLiab       string `json:"uplLiab"`
		} `json:"details"`
		Imr         string `json:"imr"`
		IsoEq       string `json:"isoEq"`
		MgnRatio    string `json:"mgnRatio"`
		Mmr         string `json:"mmr"`
		NotionalUsd string `json:"notionalUsd"`
		OrdFroz     string `json:"ordFroz"`
		TotalEq     string `json:"totalEq"`
		UTime       string `json:"uTime"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type Opt struct {
	ApiKey     string
	Passphrase string
	SecretKey  string
}

func Start(opt Opt) (string, error) {
	reqData := fetch.NewRest(fetch.NewOkxFetchOpt{
		ApiKey:     opt.ApiKey,
		Passphrase: opt.Passphrase,
		SecretKey:  opt.SecretKey,
		Path:       "/api/v5/account/balance",
		Method:     "GET",
		Data: map[string]any{
			"ccy": "USDT",
		},
	})

	if len(reqData) < 5 {
		errStr := fmt.Errorf("okx 接口请求失败")
		return "", errStr
	}

	var data ReqType
	err := jsoniter.Unmarshal(reqData, &data)
	if err != nil {
		errStr := fmt.Errorf("HotList 数据格式化失败 : " + mStr.ToStr(reqData))
		global.LogErr(mStr.ToStr(errStr))
		return "", errStr
	}

	// IP 无效也算是验证通过.其余的一概不成功
	if data.Code == "0" || data.Code == "50110" {
		balance := "0"

		if len(data.Data) > 0 {
			balance = data.Data[0].TotalEq
		}

		return balance, nil
	} else {
		errStr := fmt.Errorf(data.Msg)
		return "", errStr
	}
}
