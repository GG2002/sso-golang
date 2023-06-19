package utils

import (
	"errors"
	"log"
	"sso-golang/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var gtokenSecret []byte = []byte("00000000111111112222222233333333")

type GToken struct {
	UserName string `json:"unm"` // 用户名
	jwt.RegisteredClaims
}

func GenerateGToken(gt *GToken) string {
	gt.ExpiresAt = jwt.NewNumericDate(time.Now().Add(config.Cfg.TokenTimeout))
	// gt.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second))
	gt.IssuedAt = jwt.NewNumericDate(time.Now())
	gt.NotBefore = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, gt)

	tokenString, err := token.SignedString([]byte(gtokenSecret))

	if err != nil {
		log.Panic(err)
	}

	return tokenString
}

func CheckGToken(gtstr string) (*GToken, error) {
	token, err := jwt.ParseWithClaims(gtstr, &GToken{}, func(token *jwt.Token) (interface{}, error) {
		return gtokenSecret, nil
	})

	if claims, ok := token.Claims.(*GToken); ok {
		// fmt.Println(claims)
		if token.Valid {
			return claims, nil
		} else {
			return nil, errors.New("token is expired")
		}
	} else {
		log.Println(err)
		return nil, err
	}
}
