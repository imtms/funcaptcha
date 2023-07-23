package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/funcaptcha"
)

var url string
var ok bool

func init() {
	if url, ok = os.LookupEnv("BX_URL"); !ok {
		url = "https://bx.tms.im/202307"
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resp, err := http.Get(url)
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
