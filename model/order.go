package model

import "github.com/jinzhu/gorm"

type Order struct {//订单
	gorm.Model
	Add string//用户地址
	UserId int//购买者对应的id
	MenuID int//菜单id
	Orderserve string//服务内容
	State string//服务状态
	Orderprice int//服务金额
	Message string//用户下单备注
	Server string//服务员名称
	Tel string//服务员电话
	Ordertime string//预约时间
	Evaluate string//评价
}
