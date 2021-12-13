package routers

import (
	v1 "mosch/controller/api/v1"
	"mosch/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", v1.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", v1.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", v1.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", v1.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", v1.DeleteATodo)
	}
	return r
}
