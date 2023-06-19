package logi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"sso-golang/config"
	"sso-golang/model"
	"sso-golang/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	bodyhex, _ := io.ReadAll(c.Request.Body)
	// body := string(bodyhex)
	// log.Println(body)

	var tmp, userData model.UserData
	json.Unmarshal(bodyhex, &tmp)

	db := model.GetDBInstance()
	uNameExisted := db.CheckUserNameNum(tmp.User_name)
	if !uNameExisted {
		c.JSON(338, nil)
		return
	}
	log.Println("user:", tmp.User_name, "exists.")
	db.GetUserDataByUserName(tmp.User_name, &userData)

	m5Pwd := md5.Sum([]byte(tmp.Password + userData.Salt))
	if hex.EncodeToString(m5Pwd[:]) != userData.Password {
		c.JSON(339, nil)
		return
	}
	log.Println("user:", tmp.User_name, "log in successfully.")

	uGToken := utils.GToken{
		UserName: userData.User_name,
	}
	uGTokenString := utils.GenerateGToken(&uGToken)
	db.Set(userData.User_name, uGTokenString, config.Cfg.TokenTimeout)

	c.Header("Redirect", c.Request.Header.Get("Redirect"))
	c.JSON(302, gin.H{
		"token": uGTokenString,
	})
}

func Logout(c *gin.Context) {
	db := model.GetDBInstance()
	token, _ := c.Request.Cookie("sso_token")
	tokenClaims, _ := utils.CheckGToken(token.Value)
	db.Remove(tokenClaims.UserName)
}

func Logcheck(c *gin.Context) {
	token, _ := c.Request.Cookie("sso_token")
	tokenClaims, err := utils.CheckGToken(token.Value)
	if err != nil {
		if err.Error() == "token is expired" {
			c.JSON(340, nil)
		} else {
			c.JSON(200, gin.H{
				"Error": err.Error(),
			})
		}
		return
	}

	db := model.GetDBInstance()
	// fmt.Println(token.Value)
	if token.Value != db.Get(tokenClaims.UserName) {
		c.JSON(341, nil)
		return
	}
	c.JSON(200, "Token is correct")
}
