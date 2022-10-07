package testOkxApi

import (
	"fmt"

	"DataCenter.net/server/global"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type ReqType struct {
	Code string `bson:"code"`
	Data []struct {
		AdjEq   string `bson:"adjEq"`
		Details []struct {
			AvailBal      string `bson:"availBal"`
			AvailEq       string `bson:"availEq"`
			CashBal       string `bson:"cashBal"`
			Ccy           string `bson:"ccy"`
			CrossLiab     string `bson:"crossLiab"`
			DisEq         string `bson:"disEq"`
			Eq            string `bson:"eq"`
			EqUsd         string `bson:"eqUsd"`
			FrozenBal     string `bson:"frozenBal"`
			Interest      string `bson:"interest"`
			IsoEq         string `bson:"isoEq"`
			IsoLiab       string `bson:"isoLiab"`
			IsoUpl        string `bson:"isoUpl"`
			Liab          string `bson:"liab"`
			MaxLoan       string `bson:"maxLoan"`
			MgnRatio      string `bson:"mgnRatio"`
			NotionalLever string `bson:"notionalLever"`
			OrdFrozen     string `bson:"ordFrozen"`
			StgyEq        string `bson:"stgyEq"`
			Twap          string `bson:"twap"`
			UTime         string `bson:"uTime"`
			Upl           string `bson:"upl"`
			UplLiab       string `bson:"uplLiab"`
		} `bson:"details"`
		Imr         string `bson:"imr"`
		IsoEq       string `bson:"isoEq"`
		MgnRatio    string `bson:"mgnRatio"`
		Mmr         string `bson:"mmr"`
		NotionalUsd string `bson:"notionalUsd"`
		OrdFroz     string `bson:"ordFroz"`
		TotalEq     string `bson:"totalEq"`
		UTime       string `bson:"uTime"`
	} `bson:"data"`
	Msg string `bson:"msg"`
}

type Opt struct {
	ApiKey     string
	Passphrase string
	SecretKey  string
}

func Start(opt Opt) (string, error) {
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/account/balance",
		Data: map[string]any{
			"ccy": "USDT",
		},
		Method: "get",
		Event: func(s string, a any) {
			global.Log.Println("Event", s, a)
		},
		OKXKey: mOKX.TypeOkxKey{
			ApiKey:     opt.ApiKey,
			Passphrase: opt.Passphrase,
			SecretKey:  opt.SecretKey,
		},
	})
	if err != nil {
		errStr := fmt.Errorf("okx 接口请求失败:%+v", err)
		return "", errStr
	}

	var data ReqType
	err = jsoniter.Unmarshal(resData, &data)
	if err != nil {
		errStr := fmt.Errorf("HotList 数据格式化失败 : " + mStr.ToStr(resData))
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
