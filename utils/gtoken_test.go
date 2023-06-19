package utils

import (
	"fmt"
	"testing"
)

func TestGenerateGToken(t *testing.T) {
	tokenString := GenerateGToken(&GToken{
		UserName: "test",
	})
	fmt.Println(tokenString)
}
