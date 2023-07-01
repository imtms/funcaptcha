package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/funcaptcha"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		token, _ := funcaptcha.GetOpenAIToken()
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	r.Run(":3610")
}
