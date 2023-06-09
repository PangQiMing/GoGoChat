package main

import (
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/controller"
	"github.com/PangQiMing/GoGoChat/middleware"
	"github.com/PangQiMing/GoGoChat/socket"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//初始化Gin默认引擎
	r := gin.Default()

	//初始化数据库
	config.InitDBConfig()

	//设置static文件夹
	r.Static("/images", "./static/images")

	//启动ws服务
	hub := socket.NewHub()
	go hub.Run()

	//关闭数据库连接
	defer config.CloseDBConnection(config.DB)
	r.Use(middleware.CORSMiddleware())

	r.POST("register", controller.RegisterUser) //注册用户
	r.POST("login", controller.LoginUser)       //登录用户
	userRouters := r.Group("/user", middleware.AuthorizeJWT())
	{
		userRouters.GET("", controller.GetUserInfo)                   //获取用户个人信息
		userRouters.PUT("update-user", controller.UpdateUser)         //更新用户信息：昵称，性别，年龄
		userRouters.PUT("update-password", controller.UpdatePassword) //更新用户密码
		userRouters.POST("update-avatar", controller.UpdateAvatar)    //更新用户头像
		userRouters.GET("get-message", controller.GetMessageByGoGoID) //获取聊天记录
	}

	friendRouters := r.Group("/friend", middleware.AuthorizeJWT())
	{
		friendRouters.GET("", controller.GetFriendList)                    //获取好友列表
		friendRouters.POST("search", controller.GetSearchFriend)           //获取好友信息
		friendRouters.GET("request-list", controller.GetFriendRequestList) //好友申请列表
		friendRouters.POST("add", controller.AddFriend)                    //添加好友
		friendRouters.POST("accept", controller.AcceptFriendRequest)       //同意好友请求
		friendRouters.POST("reject", controller.RejectFriendRequest)       //拒绝好友请求
		friendRouters.POST("delete", controller.DeleteFriend)              //删除好友
	}

	groupRouters := r.Group("/group", middleware.AuthorizeJWT())
	{
		groupRouters.POST("create", controller.CreateGroup)            //创建群组
		groupRouters.PUT("update", controller.UpdateGroup)             //更新群组信息
		groupRouters.POST("delete", controller.DeleteGroup)            //解散群组
		groupRouters.POST("exit", controller.DeleteGroupMember)        //群成员退出群组
		groupRouters.POST("search", controller.GetSearchGroup)         //获取群组信息
		groupRouters.POST("join", controller.JoinGroup)                //加入群组
		groupRouters.POST("accept", controller.AcceptJoinGroupRequest) //同意入群申请
		groupRouters.POST("reject", controller.RejectJoinGroupRequest) //拒绝入群申请
		groupRouters.GET("", controller.GetGroupLists)                 //获取群组列表
		groupRouters.GET("join-list", controller.JoinGroupRequestList) //获取入群申请列表
	}

	circleRouters := r.Group("/friend-circle", middleware.AuthorizeJWT())
	{
		circleRouters.POST("add", controller.CreateFriendCircle)
		circleRouters.DELETE("delete", controller.DeleteFriendCircle)
		circleRouters.GET("", controller.GetFriendCircleAll)
	}

	r.GET("ws", func(ctx *gin.Context) {
		goGoID := ctx.Query("go_go_id")
		log.Println(goGoID)
		socket.ServeWs1(hub, goGoID, ctx.Writer, ctx.Request)
	})
	err := r.Run(":8081")
	if err != nil {
		panic("GoGoChat启动失败")
	}
}
