package main

import (
	"Graduation_project/Controller"
	"Graduation_project/DB"
	"Graduation_project/Middleware"
	"Graduation_project/model"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "regexp"
)
var db = DB.InitDB()

func main() {
        r:=gin.Default()
	  gob.Register(model.User{}) // 注册User结构体
	store := cookie.NewStore([]byte("zhr的秘密"))
	store.Options(sessions.Options{
		MaxAge: 60*120,
	})
		r.LoadHTMLGlob("templates/**/*")
	    r.Static("/static","./static")
		r.Use(Middleware.TimeJudge())
	r.Use(Middleware.CardCheck())
	r.Use(sessions.Sessions("SESSIONID", store))
	   defer db.Close()
	  db.AutoMigrate(&model.User{})
	  db.AutoMigrate(&model.Menu{})
	  db.AutoMigrate(&model.Order{})
	  db.AutoMigrate(&model.Card{})
	  db.AutoMigrate(&model.Mcard{})
	db.AutoMigrate(&model.Server{})
		r.GET("",Controller.Index())//主页面
		r.POST("/O2O",Controller.Login())
	   r.GET("/register", Controller.Toregister())//注册页面跳转
	  r.GET("/O2O", Middleware.CheckIdentity(),Controller.DetailtoIndex())//详情返回主页面
	  r.POST("/register", Controller.Register())//注册
	   r.POST("/address", Middleware.CheckIdentity(),Controller.UpdateAddress())//更新地址
	   r.GET("/update",Middleware.CheckIdentity(), Controller.ToUpdate())//跳转到更新页面
	   r.POST("/update",Middleware.CheckIdentity(),Controller.Updatepwd())//更新密码
        r.GET("/card", Middleware.CheckIdentity(),Controller.MoreCard())//更多服务卡
	  r.GET("/allmenu", Middleware.CheckIdentity(),Controller.Allmenu())//查看全部服务菜单
	  r.POST("/allmenu",Middleware.CheckIdentity(), Controller.Menusearch())//菜单页面查询
	  r.GET("/:serve",Middleware.CheckIdentity(),Middleware.HaveCard(),Controller.Todetail())//跳转到商品详情
	r.GET("/allmenu/:serve",Middleware.CheckIdentity(),Middleware.HaveCard(),Controller.Menutodetail())//菜单跳转到商品详情
	  r.GET("/card/use/:id",Middleware.CheckIdentity(), Controller.ToCardpurchase())//跳转到服务卡购买页面
	  r.POST("/card/use/:id", Middleware.CheckIdentity(),Controller.Cardpurchase())//服务卡购买逻辑
	r.GET("/:serve/card",Middleware.CheckIdentity(),Middleware.HaveCard(), Controller.To_Cardpurchase())//detail页面跳转到服务卡购买服务页面
	r.POST("/:serve/card",Middleware.CheckIdentity(),Middleware.HaveCard(),Controller.Cardpurchase2())//detail页面服务卡购买服务逻辑
	  r.GET("/:serve/purchase",Middleware.CheckIdentity(), Controller.ToServepurchase())//跳转到服务购买页面
	  r.POST("/:serve/purchase",Middleware.CheckIdentity(), Controller.Servepurchase())//服务购买逻辑
	  r.GET("/order",Middleware.CheckIdentity(),Middleware.CheckOrder(),Controller.Myorder())//我的订单
	  r.POST("/order",Middleware.CheckIdentity(), Controller.SearchOrder())//查找订单
	  r.GET("/order/:id",Middleware.CheckIdentity(), Controller.CancelOrder())//取消订单
	   r.GET("/finish/:id",Middleware.CheckIdentity(), Controller.CheckOrder())//确认完成订单
	   r.POST("/evaluate/:id",Middleware.CheckIdentity(), Controller.Evaluate())//评价
	  r.GET("/card/purchase/:id", Middleware.CheckIdentity(),Controller.ToCardpurchase2())//服务卡跳转服务卡购买页面
	  r.POST("/card/purchase/:id",Middleware.CheckIdentity(), Controller.CardPurchase3())//服务卡购买逻辑2
      r.GET("/card/buy/:id",Middleware.CheckIdentity(),Controller.CardBuy())//更多服务卡-购买月卡等
	r.POST("/card/buy/:id", Middleware.CheckIdentity(),Controller.CardPurchase2())//更多服务卡-购买月卡逻辑
	r.GET("/delete/:id",Middleware.CheckIdentity(),Controller.DeleteOrder())//删除已完成订单
	back:=r.Group("/back")
	  {
		  back.GET("/tables", Controller.ServerPage())//服务员管理界面
		  back.POST("/tables",Controller.ServerSelect())//查看服务员
		  back.POST("/tables/add",Controller.AddServer())//添加服务员
		  back.GET("/tables/id",Controller.DeleteServer())//删除服务员
		  back.POST("/tables/update",Controller.UpdateServer())//更新用户
		  back.GET("/index", func(c *gin.Context) {
			  c.HTML(200,"back-index.html",nil)
		  })
		  back.GET("/charts", Controller.Untreated())//查看待配送订单
		  back.POST("/charts/:id", Controller.Configure())//订单配置逻辑
		  back.GET("/forms", func(c *gin.Context) {
			  c.HTML(200,"forms.html",nil)
		  })
		  back.POST("/forms", Controller.AddMenu())//上传菜单项目
		  back.POST("/forms/card", Controller.AddCard())//添加服务卡
	  }
		r.Run(":8080")
}




