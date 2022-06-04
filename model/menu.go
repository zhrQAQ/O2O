package model

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {//菜单
	gorm.Model
	Servetype string`form:"servetype"`//服务类型
	Serve string`form:"serve"`//服务内容
	Introduce string`form:"introduce"`//服务介绍
	Price int`form:"price"`//单次价格
	Img string
}