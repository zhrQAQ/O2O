package model

import "github.com/jinzhu/gorm"

type Server struct {
	gorm.Model
	ServerName string`form:"name"`
	ServerState string`form:"state"`
	Age int`form:"age"`
	ServerPhone string`form:"phone"`
}
