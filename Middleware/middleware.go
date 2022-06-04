package Middleware

import (
	"Graduation_project/DB"
	"Graduation_project/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)


var timeTemplates = []string {
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}


func getCurrentUser(c *gin.Context) (userInfo model.User) {
	session := sessions.Default(c)
	userInfo = session.Get("currentUser").(model.User) // 类型转换一下
	return
}

var db = DB.InitDB()
func HaveCard()gin.HandlerFunc  {//判断是否拥有服务卡
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		cardlist:=[]model.Card{}
		serve:=c.Param("serve")
		db.Where("cardname=? AND user_id=?",serve,user.ID).Find(&cardlist)
		if len(cardlist)==0 {
			menu:=model.Menu{}
			orderlist:=[]model.Order{}
			db.Where("serve=?",serve).Find(&menu)
			db.Where("menu_id=?",menu.ID).Find(&orderlist)
			c.HTML(200,"detail.html",gin.H{
				"menu":menu,
				"user":user,
				"orderlist":orderlist,
				"error":"您未有该类的服务卡！",
			})
		}else {
			map1:= map[string]int{}
			num:=len(cardlist)
			cardtype:=[]string{}
			for i:=0;i<num;i++{//获取所有的
				m,ok:=map1[cardlist[i].Type]
				if ok {
					map1[cardlist[i].Type]=m+cardlist[i].Lastcounts
				}else  {
					map1[cardlist[i].Type]=cardlist[i].Lastcounts
					cardtype= append(cardtype, cardlist[i].Type)//将类型装入cardname
				}
			}
			for j:=0;j<len(cardtype);j++ {
				c.Set("map",map1)
				c.Set(cardtype[j],map1[cardlist[j].Type])//key：类型 value：类型对应的剩余次数
			}
			c.Next()
		}
	}
}
func TimeJudge()gin.HandlerFunc  {//判断服务卡到期时间，设置为全局中间件
	return func(c *gin.Context) {
		card:=[]model.Card{}
		db.Find(&card)
		now:=time.Now()
		num:=len(card)
		for i := 0; i < num; i++ {
			if card[i].Date.Before(now) {
				db.Delete(&card[i])
			}
		}
		c.Next()
	}
}
func CardCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		card:=[]model.Card{}
		db.Find(&card)
		for i := 0; i < len(card); i++ {
			if card[i].Lastcounts==0 {
				db.Delete(&card[i])
			}
		}
		c.Next()
	}
}
func CheckIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		if user.Name=="" {
			c.JSON(200,gin.H{
				"error":"用户信息失效",
			})
		}else {
			c.Next()
		}
	}
}
func CheckOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderlist:=[]model.Order{}
		db.Where("state=?","服务开始").Find(&orderlist)
		for i := 0; i < len(orderlist); i++ {
			t:=TimeStringToGoTime(orderlist[i].Ordertime)
			if t.After(time.Now()) {
				db.Delete(orderlist[i])
			}
		}
		c.Next()
	}
}

func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() { return t }
	}
	return time.Time{}
}
