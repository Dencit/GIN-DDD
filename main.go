package main

import (
	"app/base/config"
	Demo "app/port/demo"
	User "app/port/user"
	"github.com/gin-gonic/gin"
)

func main() {
	//路由日志-开关
	if config.Debug || config.RouteDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.Default()

	ips := []string{"127.0.0.1"}
	Router.SetTrustedProxies(ips)

	//Demo模块-路由集合
	Demo.Route(Router)
	//User模块-路由集合
	User.Route(Router)

	Router.Run(":8888")
}
