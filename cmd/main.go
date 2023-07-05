package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/funcaptcha"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resp, err := http.Get("https://bx.tms.im/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "bx error",
			})
		}
		bx, _ := io.ReadAll(resp.Body)
		token, _ := funcaptcha.GetOpenAITokenWithBx(string(bx))
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	r.Run(":3610")
}
