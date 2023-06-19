package utils

import (
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	strr := "{\"username\":2,\"password\":\"2\"}"
	strE := RsaEncrypt(strr)
	fmt.Println(strE)
	strD := RsaDecrypt(strE)
	fmt.Println(strD)
	if strD != strr {
		t.Error("解密结果与原字符串不符")
	}

	strE = RsaEncrypt(strr)
	fmt.Println(strE)
	strD = RsaDecrypt(strE)
	fmt.Println(strD)
	if strD != strr {
		t.Error("解密结果与原字符串不符")
	}
}
