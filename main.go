package main

// import "github.com/tbwisk/gormt/data/cmd"
import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tbwisk/gormt/data/view/gtools"
	"github.com/tbwisk/gormt/data/view/model"
	"github.com/tbwisk/gormt/dto"
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
		input := dto.IndexInput{}
		err := input.BindingValidParams(c)
		if err != nil {
			fmt.Println("err", err)
		}
		URL := input.Format1()
		modeldb := gtools.GetModel()
		pkg := modeldb.GenModelNew(URL)
		pkg.PackageName = input.PackageName
		title := model.Generate(pkg)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": template.HTML(title),
		})
	})
	r.Run()
}
