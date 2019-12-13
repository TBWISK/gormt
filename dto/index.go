package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// "gopkg.in/go-playground/validator.v9"

//IndexInput 首页index
type IndexInput struct {
	PackageName string `form:"package" json:"package" validate:"required"`
	Host        string `form:"host" json:"host" validate:"required"`
	DBName      string `form:"dbname" json:"dbname" validate:"required"`
	UserName    string `form:"user" json:"user" validate:"required"`
	PassWord    string `form:"password" json:"password" `
	Port        string `form:"port" json:"port" `
}

//Format1 格式转换
func (o *IndexInput) Format1() string {
	title := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		o.UserName,
		o.PassWord,
		o.Host,
		o.Port,
		o.DBName,
	)
	return title
}

//BindingValidParams 绑定校验参数
func (o *IndexInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(o); err != nil {
		return err
	}
	return nil
}
