package gtools

import (
	"github.com/tbwisk/gormt/data/view/model"
	"github.com/tbwisk/gormt/data/view/model/genmysql"
)

// GetModel get model interface. 获取model接口
func GetModel() model.IModel {
	//now just support mysql
	return &genmysql.MySQLModel
}
