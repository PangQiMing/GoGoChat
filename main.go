package main

import (
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/controller"
	"github.com/PangQiMing/GoGoChat/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//初始化Gin默认引擎
	r := gin.Default()
	//初始化数据库
	config.InitDBConfig()
	//关闭数据库连接
	defer config.CloseDBConnection(config.DB)

	userRouters := r.Group("/", middleware.AuthorizeJWT())
	{
		userRouters.POST("register", controller.RegisterUser)
	}
	err := r.Run(":8080")
	if err != nil {
		panic("GoGoChat启动失败")
	}
	log.Println("GoGoChat服务已经启动")
}
