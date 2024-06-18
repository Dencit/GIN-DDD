package demo

import (
	"app/base/midware"
	"app/port/demo/controller"
	"github.com/gin-gonic/gin"
)

// 模块路由

func Route(Router *gin.Engine) {

	SampleController := &controller.SampleController{}

	//开放路由
	Group := Router.Group("demo")
	Group.POST("sample/save", SampleController.Save)
	Group.POST("sample/update/:id", SampleController.Update)
	Group.POST("sample/delete/:id", SampleController.Delete)
	Group.GET("sample/read/:id", SampleController.Read)
	Group.GET("sample/index", SampleController.Index)

	//客户端授权路由
	UsrGroup := Router.Group("demo")
	UsrGroup.Use(midware.UsrAuth())
	UsrGroup.GET("sample/read/:id/usr", SampleController.Read)

}
