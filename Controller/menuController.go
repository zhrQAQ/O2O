package Controller

import (
	"Graduation_project/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Allmenu()gin.HandlerFunc  {//查看全部服务菜单
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		menulist:=[]model.Menu{}
		db.Find(&menulist)
		map1:= map[string]interface{}{
			user.Phone: menulist,
		}
		c.HTML(200,"allmenu.html",gin.H{
			"map1":map1,
		})
	}
}
func Menusearch()gin.HandlerFunc  {//菜单页面查询
	return func(c *gin.Context) {
		serve:=c.PostForm("serve")
		user := getCurrentUser(c)
		menulist:=[]model.Menu{}
		db.Where("serve LIKE ?","%"+serve+"%").Find(&menulist)
		map1:= map[string]interface{}{
			user.Phone: menulist,
		}
		c.HTML(200,"allmenu.html",gin.H{
			"map1":map1,
		})
	}
}
func Todetail() gin.HandlerFunc {//跳转到服务详情
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		serve:=c.Param("serve")
		menu:=model.Menu{}
		orderlist:=[]model.Order{}
		db.Where("serve=?",serve).Find(&menu)
		db.Unscoped().Where("menu_id=? AND state=?",menu.ID,"服务结束").Find(&orderlist)
		c.HTML(200,"detail.html",gin.H{
			"menu":menu,
			"user":user,
			"orderlist":orderlist,
		})
	}
}
func Menutodetail() gin.HandlerFunc {//菜单跳转到详情
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		serve:=c.Param("serve")
		menu:=model.Menu{}
		orderlist:=[]model.Order{}
		db.Where("serve=?",serve).Find(&menu)
		db.Unscoped().Where("menu_id=? AND state=?",menu.ID,"服务结束").Find(&orderlist)
		fmt.Println(orderlist)
		c.HTML(200,"detail.html",gin.H{
			"menu":menu,
			"user":user,
			"orderlist":orderlist,
		})
	}
}
func ToServepurchase() gin.HandlerFunc {//跳转到购买页面
	return func(c *gin.Context) {
		user:= getCurrentUser(c)
		serve:=c.Param("serve")
		menu:=model.Menu{}
		db.Where("serve=?",serve).Find(&menu)
		c.HTML(200,"purchase.html",gin.H{
			"user":user,
			"menu":menu,
		})
	}
}
func Servepurchase() gin.HandlerFunc {//服务购买逻辑
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		serve:=c.Param("serve")
		msg:=c.PostForm("msg")
		timestring:=c.PostForm("time")
		menu:=model.Menu{}
		db.Where("serve=?",serve).Find(&menu)
		order:=model.Order{Add:user.Address,Orderprice:menu.Price,UserId: int(user.ID),Orderserve:serve,MenuID: int(menu.ID),Message: msg,State: "服务开始",Ordertime: timestring}
		db.Create(&order)
		c.HTML(200,"purchase.html",gin.H{
			"user":user,
			"menu":menu,
			"success":"购买成功！",
		})
	}
}
func DetailtoIndex() gin.HandlerFunc {//详情返回主页
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		cardlist:=[]model.Card{}
		db.Where("user_id=?",user.ID).Find(&cardlist)
		c.HTML(200,"index.html",gin.H{
			"cardlist":cardlist,
			"user":user,
		})
	}
}
