package model

import "github.com/jinzhu/gorm"

type Mcard struct {//服务卡
	gorm.Model
	Mcardname string//服务卡名称
	Mcardtype string//服务卡类型，月卡、季卡、年卡
	Mcardcount int//服务卡服务次数
	Mcardprice int//服务卡价格
}
