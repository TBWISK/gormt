package main

// import "github.com/xxjwxc/gormt/data/cmd"
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// cmd.Execute()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.POST("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run()
}
