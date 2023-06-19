package model

import (
	"fmt"
	"sso-golang/config"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println("f")
	db := GetDBInstance()
	db.Set("i", "k", config.Cfg.TokenTimeout)
	fmt.Println(db.Get("i"))
}
