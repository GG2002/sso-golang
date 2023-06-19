package sign

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"sso-golang/model"
	"time"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	bodyhex, _ := io.ReadAll(c.Request.Body)
	// body := string(bodyhex)
	// log.Println(body)

	var userData model.UserData
	json.Unmarshal(bodyhex, &userData)
	userData.Salt = "oneecho" + time.Now().Format("2006-01-02 15:04:05")
	// log.Println(userData)

	m5Pwd := md5.Sum([]byte(userData.Password + userData.Salt))
	userData.Password = hex.EncodeToString(m5Pwd[:])
	// log.Println(userData.Password)

	db := model.GetDBInstance()
	unameDup := db.CheckUserNameNum(userData.User_name)
	if unameDup {
		// 用户名重复
		c.JSON(337, nil)
		return
	}
	err := db.InsertUserData(&userData)
	if err != nil {
		c.JSON(200, gin.H{
			"Error": err,
		})
	}
}

func UsernameExisted(c *gin.Context) {
	bodyhex, _ := io.ReadAll(c.Request.Body)
	// body := string(bodyhex)
	// log.Println(body)

	var userData model.UserData
	json.Unmarshal(bodyhex, &userData)

	db := model.GetDBInstance()
	unameDup := db.CheckUserNameNum(userData.User_name)

	if unameDup {
		// 用户名重复
		c.JSON(337, nil)
	} else {
		c.JSON(200, nil)
	}
}
