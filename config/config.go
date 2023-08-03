package config

import (
	"os"
	"time"
)

var Cfg struct {
	RSAPubKeyGBase string
	RSAPriKeyGBase string
	TokenTimeout   time.Duration
}

func init() {
	curDir, _ := os.Getwd()
	Cfg.RSAPubKeyGBase = curDir + "/utils/public.pem"
	Cfg.RSAPriKeyGBase = curDir + "/utils/private.pem"
	Cfg.TokenTimeout = 24 * time.Hour
}
