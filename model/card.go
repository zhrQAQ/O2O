package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Card struct {//用户服务卡结构体
	gorm.Model
	UserID int//持有者所对应的号码
	Type string//服务卡类型
	Cardname string//服务卡名称
	Date time.Time//到期时间
	Lastcounts int//剩余次数
	Usedcounts int //已使用次数
	Cardprice int//服务卡购买价格
}
