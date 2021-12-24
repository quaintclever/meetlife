package main

import (
	"fmt"
	"mosch/dao"
	"mosch/models"
	"mosch/routers"
	"mosch/setting"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./mosch config/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Println("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Println("server startup failed, err:%v\n", err)
	}
}