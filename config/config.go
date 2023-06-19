package config

import "time"

var Cfg struct {
	RSAPubKeyGBase string
	RSAPriKeyGBase string
	TokenTimeout   time.Duration
}

func init() {
	Cfg.RSAPubKeyGBase = "C:/Users/77016/Desktop/HustmathsRecruit/sso-golang/utils/public.pem"
	Cfg.RSAPriKeyGBase = "C:/Users/77016/Desktop/HustmathsRecruit/sso-golang/utils/private.pem"
	Cfg.TokenTimeout = 24 * time.Hour
}
