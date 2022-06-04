package Controller

import (
	"Graduation_project/DB"
	"Graduation_project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var db = DB.InitDB()

func MoreCard()gin.HandlerFunc  {//查看更多服务卡
	return func(c *gin.Context) {
		monthcard:=[]model.Mcard{}
		quartercard:=[]model.Mcard{}
		yearcard:=[]model.Mcard{}
		db.Where("mcardtype=?","月卡").Find(&monthcard)
		db.Where("mcardtype=?","季卡").Find(&quartercard)
		db.Where("mcardtype=?","年卡").Find(&yearcard)
		fmt.Println(monthcard)
		c.HTML(200,"card.html",gin.H{
			"monthcard":monthcard,
			"quartercard":quartercard,
			"yearcard":yearcard,
		})
	}
}
func ToCardpurchase() gin.HandlerFunc {//跳转到服务卡购买页面
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		id:=c.Param("id")
		card:=model.Card{}
		db.Where("id=?",id).Find(&card)
		menu:=model.Menu{}
		db.Where("serve=?",card.Cardname).Find(&menu)
		c.HTML(200,"card_purchase2.html",gin.H{
			"user":user,
			"menu":menu,
		})
	}
}
func To_Cardpurchase() gin.HandlerFunc {//详情跳转购买
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		map1,ok:=c.Get("map")//获取数量
		if !ok{
			c.JSON(200,gin.H{
				"msg":ok,
			})
		}
		fmt.Println("接受后的card:",map1)
		serve:=c.Param("serve")
		menu:=model.Menu{}
		db.Where("serve=?",serve).Find(&menu)
		c.HTML(200,"card_purchase.html",gin.H{
			"user":user,
			"menu":menu,
			"map1":map1,
		})
	}
}
func ToCardpurchase2() gin.HandlerFunc {//跳转服务卡购买页面
	return func(c *gin.Context) {
		id:=c.Param("id")
		card:=model.Card{}
		db.Where("id=?",id).Find(&card)
		mcard:=model.Mcard{}
		db.Where("mcardname=? AND mcardtype=?",card.Cardname,card.Type).Find(&mcard)
		c.HTML(200,"cardpurchase.html",gin.H{
			"mcard":mcard,
		})
	}
}
func Cardpurchase() gin.HandlerFunc {//服务卡购买逻辑
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		id:=c.Param("id")
		msg:=c.PostForm("msg")
		timestring:=c.PostForm("time")
		card:=model.Card{}
		db.Where("id=?",id).Find(&card)
		menu:=model.Menu{}
		db.Where("serve=? ",card.Cardname).Find(&menu)
		order:=model.Order{
			Add:user.Address,
			Orderprice:menu.Price,
			UserId: int(user.ID),
			Orderserve:card.Cardname,
			MenuID: int(menu.ID),
			Message: msg,
			State: "服务开始",
			Ordertime: timestring}
		db.Create(&order)//生成订单
		db.Model(&card).Update(map[string]interface{}{"lastcounts":card.Lastcounts-1,"usedcounts":card.Usedcounts+1})
			c.HTML(200,"card_purchase2.html",gin.H{
				"user":user,
				"menu":menu,
				"success":"购买成功！",
			})
	}
}
func Cardpurchase2() gin.HandlerFunc {//服务卡购买逻辑
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		serve:=c.Param("serve")
		msg:=c.PostForm("msg")
		t:=c.PostForm("cardtype")
		timestring:=c.PostForm("time")
		cardlist:=[]model.Card{}
		db.Where("cardname=? AND type=? AND user_id=?",serve,t,user.ID).Find(&cardlist)
		card:=model.Card{}
		if len(cardlist)>1 {
			for i := 0; i < len(cardlist)-1; i++ {
				if cardlist[i].Date.Before(cardlist[i+1].Date){
					card=cardlist[i]
				}
			}
		}else {
		     card=cardlist[0]
		}
		fmt.Println("最近的card为：",card)
		menu:=model.Menu{}
		db.Where("serve=? ",card.Cardname).Find(&menu)
		order:=model.Order{Add:user.Address,Orderprice:menu.Price,UserId: int(user.ID),Orderserve:card.Cardname,MenuID: int(menu.ID),Message: msg,State: "服务开始",Ordertime: timestring}
		db.Create(&order)//生成订单
		db.Model(&card).Update(map[string]interface{}{"lastcounts":card.Lastcounts-1,"usedcounts":card.Usedcounts+1})
		if card.Lastcounts==0 {
			db.Delete(&card)
		}else {
			map1,ok:=c.Get("map")//获取数量
			if !ok{
				c.JSON(200,gin.H{
					"msg":ok,
				})
			}
			c.HTML(200,"card_purchase.html",gin.H{
				"user":user,
				"menu":menu,
				"success":"购买成功！",
				"map1":map1,
			})
		}
	}
}
func CardPurchase2() gin.HandlerFunc {//服务卡购买逻辑
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		id:=c.Param("id")
		count:=c.PostForm("count")
		newcount,_:=strconv.Atoi(count)
		mcard:=model.Mcard{}
		db.Where("id=?",id).Find(&mcard)
		now:=time.Now()
		if mcard.Mcardtype=="月卡"{
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*30*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}else if mcard.Mcardtype=="季卡" {
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*90*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}else if mcard.Mcardtype=="年卡" {
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*365*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}
	}
}
func CardPurchase3() gin.HandlerFunc {//服务卡购买逻辑
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		id:=c.Param("id")
		count:=c.PostForm("count")
		newcount,_:=strconv.Atoi(count)
		mcard:=model.Mcard{}
		card:=model.Card{}
		db.Where("id=?",id).Find(&card)
		db.Where("mcardname=? AND mcardtype=? ",card.Cardname,card.Type).Find(&mcard)
		now:=time.Now()
		if mcard.Mcardtype=="月卡"{
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*30*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}else if mcard.Mcardtype=="季卡" {
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*90*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}else if mcard.Mcardtype=="年卡" {
			month:=time.Duration(newcount)
			dd, _ := time.ParseDuration("24h")
			dd1:=now.Add(dd*365*month)
			card:=model.Card{Cardname:mcard.Mcardname,UserID: int(user.ID),Type: mcard.Mcardtype,Date: dd1,Lastcounts: mcard.Mcardcount,Cardprice: mcard.Mcardprice*newcount}
			db.Create(&card)
			c.HTML(200,"cardpurchase.html",gin.H{
				"mcard":mcard,
				"success":"购买成功！",
			})
		}
	}
}