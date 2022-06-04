package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB()*gorm.DB  {
	db, err := gorm.Open("mysql", "root:zhr19991129@(127.0.0.1:3306)/O2O?charset=utf8mb4&parseTime=True&loc=Local")//数据库
	if err!= nil{
		panic(err)
	}
	return db
}
