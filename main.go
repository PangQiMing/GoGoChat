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

	r.POST("register", controller.RegisterUser) //注册用户
	r.POST("login", controller.LoginUser)       //登录用户
	userRouters := r.Group("/user", middleware.AuthorizeJWT())
	{
		userRouters.GET("", controller.GetUserInfo)           //获取用户个人信息
		userRouters.POST("logout", controller.LogoutUser)     //退出登录
		userRouters.PUT("update-user", controller.UpdateUser) //更新用户信息：如头像，昵称，性别，年龄
	}

	friendRouters := r.Group("/friend", middleware.AuthorizeJWT())
	{
		friendRouters.GET("", controller.GetFriendList)                    //获取好友列表
		friendRouters.GET("request-list", controller.GetFriendRequestList) //好友申请列表
		friendRouters.POST("add", controller.AddFriend)                    //添加好友
		friendRouters.POST("accept", controller.AcceptFriendRequest)       //同意好友请求
		friendRouters.POST("reject", controller.RejectFriendRequest)       //拒绝好友请求
		friendRouters.DELETE("delete", controller.DeleteFriend)            //删除好友
	}

	log.Println("GoGoChat服务启动中")
	err := r.Run(":8080")
	if err != nil {
		panic("GoGoChat启动失败")
	}
}
