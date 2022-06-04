package Controller

import (
	"Graduation_project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Untreated() gin.HandlerFunc {//查看待配送订单
	return func(c *gin.Context) {
		orderlist:=[]model.Order{}
		serverlist:=[]model.Server{}
		db.Where("state=?","服务开始").Find(&orderlist)
		db.Where("server_state=?","空闲").Find(&serverlist)
		fmt.Println(orderlist)
		c.HTML(200,"charts.html",gin.H{
			"orderlist":orderlist,
			"serverlist":serverlist,
		})
	}
}
func Configure() gin.HandlerFunc {//订单配置逻辑
	return func(c *gin.Context) {
		waiter:=c.PostForm("waiter")
		id:=c.Param("id")
		fmt.Println(waiter)
		order:=model.Order{}
		db.Where("id=?",id).Find(&order)
		server:=model.Server{}
		db.Where("server_name=?",waiter).Find(&server)
		db.Debug().Model(&order).Update(map[string]interface{}{"state":"服务进行中","server":waiter,"tel":server.ServerPhone})
		fmt.Println(server.ServerPhone)
		db.Model(&server).Update("server_state","服务中")
		orderlist:=[]model.Order{}
		serverlist:=[]model.Server{}
		db.Where("state=?","服务开始").Find(&orderlist)
		db.Where("server_state=?","空闲").Find(&serverlist)
		c.HTML(200,"charts.html",gin.H{
			"orderlist":orderlist,
			"success":"订单配置成功！",
			"serverlist":serverlist,
		})
	}
}
func AddMenu() gin.HandlerFunc {//上传菜单项目
	return func(c *gin.Context) {
		var menu model.Menu
		c.ShouldBind(&menu)
		menu1:=model.Menu{}
		db.Where("servetype=?",menu.Serve).Find(&menu1)
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		dst := fmt.Sprintf("D:/code/goland/Graduation_project/static/images/%s", file.Filename)
		if menu1.Serve==""{
			menu.Img=file.Filename
			db.Create(&menu)
			c.SaveUploadedFile(file, dst)
			c.HTML(200,"forms.html",nil)
		}else {
			c.HTML(200,"forms.html",gin.H{
				"error":"已经有该服务类型",
			})
		}
	}
}
func AddCard() gin.HandlerFunc {//添加服务卡
	return func(c *gin.Context) {
		cardname := c.PostForm("cardname") //获取服务卡名称
		cardtype := c.PostForm("cardtype") //获取服务卡类型
		count := c.PostForm("count")
		newcount, _ := strconv.Atoi(count)
		price := c.PostForm("price")
		newprice, _ := strconv.Atoi(price)
		card := model.Mcard{}
		db.Where("mcardname=? AND mcardtype=?", cardname, cardtype).Find(&card)
		if card.Mcardname == "" {
			card = model.Mcard{Mcardname: cardname, Mcardtype: cardtype, Mcardcount: newcount, Mcardprice: newprice}
			db.Create(&card)
			c.HTML(200, "forms.html", gin.H{
				"success": "添加成功！",
			})
		} else {
			c.HTML(200, "forms.html", gin.H{
				"error": "已经有该服务卡！",
			})
		}
	}
}
func ServerPage() gin.HandlerFunc {//跳转服务员页面
	return func(c *gin.Context) {
		server:=[]model.Server{}
		db.Find(&server)
		c.HTML(200,"tables.html",gin.H{
			"server":server,
		})
	}
}
func ServerSelect() gin.HandlerFunc {//查看服务员
	return func(c *gin.Context) {
		name:=c.PostForm("name")
		state:=	c.PostForm("state")
		if name==""{
			server:=[]model.Server{}
			if state!="" {
				db.Where("server_state=?",state).Find(&server)
				c.HTML(200,"tables.html",gin.H{
					"server":server,
				})
			}else {
				db.Find(&server)
				c.HTML(200,"tables.html",gin.H{
					"server":server,
				})
			}
		}else {
			    server2:=[]model.Server{}
				db.Where("server_state=? AND server_name=?",state,name).Find(&server2)
				c.HTML(200,"tables.html",gin.H{
					"server":server2,
				})
		}

	}
}
func DeleteServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		id:=c.Param("id")
		server:=model.Server{}
		db.Where("id=?",id).Find(&server)
		db.Delete(&server)
		server2:=[]model.Server{}
		db.Find(&server2)
		c.HTML(200,"tables.html",gin.H{
			"server":server2,
		})
	}
}
func UpdateServer() gin.HandlerFunc {//更新服务员
	return func(c *gin.Context) {
		id:=c.PostForm("id")
		phone:=c.PostForm("phone")
		age:=c.PostForm("age")
		newage,_:=strconv.Atoi(age)
		server:=model.Server{}
		db.Where("id=?",id).Find(&server)
		server2:=model.Server{}
		db.Where("server_phone=?",phone).Find(&server2)
		if server2.ServerPhone=="" {
			db.Model(&server).Update(map[string]interface{}{"age":newage,"server_phone":phone})
			server2:=[]model.Server{}
			db.Find(&server2)
			c.HTML(200,"tables.html",gin.H{
				"server":server2,
			})
		}else {
			c.JSON(200,gin.H{
				"error":"手机号已经存在！",
			})
		}
	}
}
func CardBuy() gin.HandlerFunc {//更多服务卡-购买月卡等
	return func(c *gin.Context) {
		id:=c.Param("id")
		mcard:=model.Mcard{}
		db.Where("id=?",id).Find(&mcard)
		c.HTML(200,"cardpurchase.html",gin.H{
			"mcard":mcard,
		})
	}
}