package tmpl

import (
	_ "embed"
)

//go:embed inst_CoinAI.sh
var InstCoinServe string

type InstCoinServeParam struct {
	Port   string
	UserID string
}
