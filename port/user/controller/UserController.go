package controller

import (
	"app/base/request"
	"app/base/respond"
	UserLogic "app/port/user/logic"
	UserQuery "app/port/user/query"
	"github.com/gin-gonic/gin"
)

/**
notes: 应用层-控制器
说明: 控制器内不写业务,只写http层面相关的逻辑,
调用原则: 向下调用[输入验证类,业务类,输出转化类].
*/

//控制器结构

type UserController struct{}

//详情

func (instance *UserController) Read(context *gin.Context) {
	id := context.Param("id")

	//业务逻辑
	result := (&UserLogic.UserLogic{}).Read(context, id)

	respond.Json(context).Read(result)
}

//列表

func (instance *UserController) Index(context *gin.Context) {
	//query
	queryStd := UserQuery.UserQuery{}
	query := request.Query(context).Get(&queryStd)

	//业务逻辑
	result, mata := (&UserLogic.UserLogic{}).Index(context, query)

	respond.Json(context).Index(result, mata)
}
