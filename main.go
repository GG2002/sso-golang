package main

import (
	"log"
	"net/http"
	"os"
	"sso-golang/logi"
	"sso-golang/sign"
	"sso-golang/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		log.Println(origin)
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "http://hustmaths.top") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Redirect")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Redirect")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})
	logFile, _ := os.Create("./sso-golang.log")
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Llongfile)

	signn := r.Group("/sign")
	signn.Use(utils.RsaDecryptMiddleWare)
	{
		signn.POST("/signup", sign.Signup)
		signn.POST("/usernameexisted", sign.UsernameExisted)
	}

	logii := r.Group("/logi")
	logii.Use(utils.RsaDecryptMiddleWare)
	{
		logii.POST("/login", logi.Login)
		logii.POST("/logout", logi.Logout)
		logii.POST("/logcheck", logi.Logcheck)
	}
	r.POST("/getpubkey", utils.GetPubKey)
	r.Run(":11451")
}
