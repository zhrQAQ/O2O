package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {//用户结构体
	gorm.Model
	Email string`form:"email"`
	Name string`form:"username"`
	Password string`form:"password""`
	Address string`form:"address"`
	Phone string`form:"phone"`
}




