package Controller

import (
	"Graduation_project/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Myorder() gin.HandlerFunc {//我的订单
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		order1:=[]model.Order{}
		order2:=[]model.Order{}
		order3:=[]model.Order{}
		db.Where("user_id=? AND state=?",user.ID,"服务开始").Find(&order1)
		db.Where("user_id=? AND state=?",user.ID,"服务进行中").Find(&order2)
		db.Where("user_id=? AND state=?",user.ID,"服务结束").Find(&order3)
		c.HTML(200,"order.html",gin.H{
			"order1":order1,
			"order2":order2,
			"order3":order3,
			"user":user,
		})
	}
}
func SearchOrder() gin.HandlerFunc {//查找订单
	return func(c *gin.Context) {
		limit:=c.PostForm("limit")
		serve:=c.PostForm("serve")
		user := getCurrentUser(c)
		if limit=="" {
			order1:=[]model.Order{}
			order2:=[]model.Order{}
			order3:=[]model.Order{}
			db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,"服务开始","%"+serve+"%").Find(&order1)
			db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,"服务进行中","%"+serve+"%").Find(&order2)
			db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,"服务结束","%"+serve+"%").Find(&order3)
			fmt.Println(limit)
			fmt.Println(order1)
			c.HTML(200,"order.html",gin.H{
				"order1":order1,
				"order2":order2,
				"order3":order3,
			})
		}else {
			order1:=[]model.Order{}
			order2:=[]model.Order{}
			order3:=[]model.Order{}
			if limit=="服务开始" {
				db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,limit,"%"+serve+"%").Find(&order1)
			}else if limit=="服务进行中" {
				db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,limit,"%"+serve+"%").Find(&order2)
			}else if limit =="服务结束"{
				db.Where("user_id=? AND state=? AND orderserve LIKE ?",user.ID,limit,"%"+serve+"%").Find(&order3)
			}
			c.HTML(200,"order.html",gin.H{
				"order1":order1,
				"order2":order2,
				"order3":order3,
			})
		}
	}
}
func CancelOrder() gin.HandlerFunc {//取消订单
	return func(c *gin.Context) {
		id:=c.Param("id")
		user := getCurrentUser(c)
		db.Debug().Where("id=?",id).Delete(&model.Order{})
		order1:=[]model.Order{}
		order2:=[]model.Order{}
		order3:=[]model.Order{}
		db.Where("user_id=? AND state=?",user.ID,"服务开始").Find(&order1)
		db.Where("user_id=? AND state=?",user.ID,"服务进行中").Find(&order2)
		db.Where("user_id=? AND state=?",user.ID,"服务结束").Find(&order3)
		c.HTML(200,"order.html",gin.H{
			"order1":order1,
			"order2":order2,
			"order3":order3,
			"success":"取消成功！",
		})
	}
}
func CheckOrder() gin.HandlerFunc {//确认完成订单
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		id:=c.Param("id")
		order:=model.Order{}
		db.Where("id=?",id).Find(&order)
		name:=order.Server
		server:=model.Server{}
		db.Where("server_name=?",name).Find(&server)
		db.Model(&server).Update("server_state","空闲")
		db.Model(&order).Update(map[string]interface{}{"state":"服务结束"})
		order1:=[]model.Order{}
		order2:=[]model.Order{}
		order3:=[]model.Order{}
		db.Where("user_id=? AND state=?",user.ID,"服务开始").Find(&order1)
		db.Where("user_id=? AND state=?",user.ID,"服务进行中").Find(&order2)
		db.Where("user_id=? AND state=?",user.ID,"服务结束").Find(&order3)
		c.HTML(200,"order.html",gin.H{
			"order1":order1,
			"order2":order2,
			"order3":order3,
		})
	}
}
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id:=c.Param("id")
		order:=model.Order{}
		db.Where("id=?",id).Find(&order)
		db.Delete(&order)
		user := getCurrentUser(c)
		order1:=[]model.Order{}
		order2:=[]model.Order{}
		order3:=[]model.Order{}
		db.Where("user_id=? AND state=?",user.ID,"服务开始").Find(&order1)
		db.Where("user_id=? AND state=?",user.ID,"服务进行中").Find(&order2)
		db.Where("user_id=? AND state=?",user.ID,"服务结束").Find(&order3)
		c.HTML(200,"order.html",gin.H{
			"order1":order1,
			"order2":order2,
			"order3":order3,
		})
	}
}
