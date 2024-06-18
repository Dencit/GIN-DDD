package demo

import (
	"app/port/user/controller"
	"github.com/gin-gonic/gin"
)

// 开放路由

func Route(Router *gin.Engine) {

	UserController := &controller.UserController{}

	//开放路由
	Group := Router.Group("user")
	Group.GET("user/read/:id", UserController.Read)
	Group.GET("user/index", UserController.Index)

	//客户端授权路由

}
