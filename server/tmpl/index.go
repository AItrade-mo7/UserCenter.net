package tmpl

import (
	_ "embed"
	"time"
)

//go:embed email-sys.html
var SysEmail string

type SysParam struct {
	Message      string
	SysTime      time.Time
	SecurityCode string
}

//go:embed email-code.html
var CodeEmail string

type CodeParam struct {
	VerifyCode   string
	Action       string
	SysTime      string
	SecurityCode string
}

//go:embed email-register.html
var RegisterSucceedEmail string

type RegisterSucceedParam struct {
	Password     string
	SysTime      string
	SecurityCode string
}

//go:embed install_CoinFund.sh
var InstCoinServe string

type InstCoinServeParam struct {
	Port        string
	UserID      string
	CoinServeID string
}
