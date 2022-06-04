package Controller

import (
	"Graduation_project/model"
	"github.com/gin-gonic/gin"
)

func AddServer() gin.HandlerFunc {//添加服务员
	return func(c *gin.Context) {
		var server model.Server
		c.ShouldBind(&server)
		s:=model.Server{}
		db.Where("server_phone=?",server.ServerPhone).Find(&s)
		if s.ServerPhone=="" {
			server.ServerState="空闲"
			db.Create(&server)
			server2:=[]model.Server{}
			db.Find(&server2)
			c.HTML(200,"tables.html",gin.H{
				"success":"添加成功！",
				"server":server2,
			})
		}else {
			server2:=[]model.Server{}
			db.Find(&server2)
			c.HTML(200,"tables.html",gin.H{
				"error":"已经存在该服务员！",
				"server":server2,
			})
		}
	}
}
