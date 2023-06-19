package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"io"
	"log"
	"os"
	"sso-golang/config"

	"github.com/gin-gonic/gin"
)

var puKey *rsa.PublicKey
var prKey *rsa.PrivateKey

// 生成Rsa密钥，若两个密钥文件都存在则使用原文件
func init() {
	var err error
	_, err1 := os.Stat(config.Cfg.RSAPriKeyGBase)
	_, err2 := os.Stat(config.Cfg.RSAPubKeyGBase)
	if err1 == nil && err2 == nil {
		prKeyBuffer, err := os.ReadFile(config.Cfg.RSAPriKeyGBase)
		if err != nil {
			log.Panic(err)
		}
		prBlock, _ := pem.Decode([]byte(prKeyBuffer))
		if prBlock == nil {
			log.Panic(err)
		}
		prKeyInterface, err := x509.ParsePKCS8PrivateKey(prBlock.Bytes)
		prKey = prKeyInterface.(*rsa.PrivateKey)
		if err != nil {
			log.Panic(err)
		}

		puKeyBuffer, err := os.ReadFile(config.Cfg.RSAPubKeyGBase)
		if err != nil {
			log.Panic(err)
		}
		puBlock, _ := pem.Decode([]byte(puKeyBuffer))
		if puBlock == nil {
			log.Panic(err)
		}
		puKeyInterface, err := x509.ParsePKIXPublicKey(puBlock.Bytes)
		puKey = puKeyInterface.(*rsa.PublicKey)
		if err != nil {
			log.Panic(err)
		}

		return
	}
	prKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panic(err)
	}
	puKey = &prKey.PublicKey

	filePr, err := os.Create(config.Cfg.RSAPriKeyGBase)
	if err != nil {
		log.Panic(err)
	}

	filePu, err := os.Create(config.Cfg.RSAPubKeyGBase)
	if err != nil {
		log.Panic(err)
	}
	prKeyBytes, _ := x509.MarshalPKCS8PrivateKey(prKey)
	err = pem.Encode(filePr, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: prKeyBytes,
	})
	if err != nil {
		log.Panic(err)
	}
	puKeyBytes, _ := x509.MarshalPKIXPublicKey(puKey)
	err = pem.Encode(filePu, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: puKeyBytes,
	})
	if err != nil {
		log.Panic(err)
	}
}

// Rsa公钥加密str
func RsaEncrypt(str string) string {
	// 加密str
	sha256hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(sha256hash, rand.Reader, puKey, []byte(str), nil)
	if err != nil {
		log.Panic(err)
	}
	encryptedText := base64.StdEncoding.EncodeToString(ciphertext)

	return encryptedText
}

// Rsa私钥解密str
func RsaDecrypt(str string) string {
	// 解密str
	decodedText, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Panic("base64 decode failed", err.Error())
	}
	sha256hash := sha256.New()
	decryptedText, err := rsa.DecryptOAEP(sha256hash, nil, prKey, decodedText, nil)
	if err != nil {
		log.Panic(err)
	}

	return string(decryptedText)
}

func GetPubKey(c *gin.Context) {
	puKeyBuffer, err := os.ReadFile(config.Cfg.RSAPubKeyGBase)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(200, gin.H{
		"rsaPubKey": string(puKeyBuffer),
	})
}

func RsaDecryptMiddleWare(c *gin.Context) {
	bodyhex, _ := io.ReadAll(c.Request.Body)
	if len(bodyhex) == 0 {
		return
	}
	bodyEncryptedBytes, _ := hex.DecodeString(string(bodyhex))
	bodyBytes, _ := hex.DecodeString(RsaDecrypt(base64.StdEncoding.EncodeToString(bodyEncryptedBytes)))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	log.Println(string(bodyBytes))
}
