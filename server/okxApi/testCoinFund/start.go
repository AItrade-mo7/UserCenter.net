package testCoinFund

import "fmt"

type Opt struct {
	ApiKey     string
	Passphrase string
	SecretKey  string
}

func Start() (string, error) {
	errStr := fmt.Errorf("接口尚在开发中")
	return "", errStr
}
