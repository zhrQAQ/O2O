package Controller

import (
	"Graduation_project/model"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
func getCurrentUser(c *gin.Context) (userInfo model.User) {
	session := sessions.Default(c)
	userInfo = session.Get("currentUser").(model.User) // 类型转换一下
	return
}
func setCurrentUser(c *gin.Context, userInfo model.User) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}
func Index() gin.HandlerFunc {//主页面
	return func(c *gin.Context) {
		c.HTML(200,"login.html",nil)
	}
}
func Toregister() gin.HandlerFunc {//注册页面跳转
	return func(c *gin.Context) {
		c.HTML(200,"register.html",nil)
	}
}
func Login()gin.HandlerFunc{//用户登录判断
	return func(c *gin.Context) {
		gob.Register(model.User{}) // 注册User结构体
		phone:=c.PostForm("phone")
		password:=c.PostForm("password")
		user:=model.User{}
		db.Debug().Where("phone=? AND password=?",phone,password).Find(&user)
		if phone=="root"&&password=="zhr19991129"{
			c.HTML(200,"back-index.html",nil)
		}else {
			if user.Phone==""{
				c.HTML(200,"login.html",gin.H{
					"error":"账号或者密码错误！",
				})
			}else {
				fmt.Println(user)
				cardlist:=[]model.Card{}
				db.Where("user_id=?",user.ID).Find(&cardlist)
				setCurrentUser(c, user) // 电话和密码正确则将当前用户信息写入session中
				c.HTML(200,"index.html",gin.H{
					"cardlist":cardlist,
					"user":user,
				})
			}
		}
	}
}
func Register() gin.HandlerFunc {//用户注册逻辑
	return func(c *gin.Context) {
		var user model.User
		err:=c.ShouldBind(&user)
		fmt.Println(user)
		if err!=nil{
			c.JSON(200,gin.H{
				"error":"无效的数据",
			})
		}else {
			user1:=model.User{}
			db.Where("phone=?",user.Phone).Find(&user1)
			if user1.Phone==user.Phone {
				c.HTML(200,"register.html",gin.H{
					"error":"已经存在该用户！",
				})
			}else {
				db.Create(&user)
				c.HTML(200,"register.html",gin.H{
					"success":"注册成功！",
				})
			}
		}
	}
}
func ToUpdate() gin.HandlerFunc {//跳转到更新页面
	return func(c *gin.Context) {
		user := getCurrentUser(c)
		c.HTML(200,"info_change.html",gin.H{
			"user":user,
		})
	}
}

func UpdateAddress() gin.HandlerFunc {//修改地址
	return func(c *gin.Context) {
		address:=c.PostForm("address")
		user := getCurrentUser(c)
		if address==user.Address {
			c.HTML(200,"info_change.html",gin.H{
				"user":user,
				"error2":"地址不能与原地址相同！",
			})
		}else {
			db.Model(&user).Update("address",address)
			c.HTML(200,"info_change.html",gin.H{
				"user":user,
				"success2":"地址修改成功！",
			})
		}
	}
}
func Updatepwd()gin.HandlerFunc  {//修改密码
	return func(c *gin.Context) {
		oldpwd:=c.PostForm("old")
		newpwd:=c.PostForm("new")
		user := getCurrentUser(c)
		if user.Password==oldpwd{
			if oldpwd!=newpwd {
				db.Model(&user).Update("password",newpwd)
				c.HTML(200,"info_change.html",gin.H{
					"user":user,
					"success":"修改成功！",
				})
			}else {
				c.HTML(200,"info_change.html",gin.H{
					"user":user,
					"error":"新密码不能与原密码相同！",
				})
			}
		}
	}
}
func Evaluate() gin.HandlerFunc {//评价
	return func(c *gin.Context) {
		evaluate:=c.PostForm("evaluate")
		id:=c.Param("id")
		order:=model.Order{}
		db.Where("id=?",id).Find(&order)
		db.Model(&order).Update("evaluate",evaluate)
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
