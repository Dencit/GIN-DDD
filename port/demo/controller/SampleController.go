package controller

import (
	"app/base/request"
	"app/base/respond"
	"app/base/validate"
	DemoLogic "app/port/demo/logic"
	DemoQuery "app/port/demo/query"
	DemoRequest "app/port/demo/request"
	"github.com/gin-gonic/gin"
)

/**
notes: 应用层-控制器
说明: 控制器内不写业务,只写http层面相关的逻辑,
调用原则: 向下调用[输入验证类,业务类,输出转化类].
*/

//控制器结构

type SampleController struct{}

//新增

func (instance *SampleController) Save(context *gin.Context) {

	//输入验证
	inputStd := DemoRequest.SampleSave{}
	input := request.Json(context).Input(&inputStd)
	validate.Check(context).Command(&inputStd)

	//业务逻辑
	result := (&DemoLogic.SampleLogic{}).Save(context, input)

	respond.Json(context).Save(result)
}

//更新

func (instance *SampleController) Update(context *gin.Context) {

	//路径ID
	id := context.Param("id")

	//输入验证
	inputStd := DemoRequest.SampleUpdate{}
	input := request.Json(context).Input(&inputStd)
	validate.Check(context).Command(&inputStd)

	//业务逻辑
	result := (&DemoLogic.SampleLogic{}).Update(context, id, input)

	respond.Json(context).Update(result)
}

//删除

func (instance *SampleController) Delete(context *gin.Context) {
	//路径ID
	id := context.Param("id")

	//业务逻辑
	(&DemoLogic.SampleLogic{}).Delete(context, id)

	respond.Json(context).Delete()
}

//详情

func (instance *SampleController) Read(context *gin.Context) {
	id := context.Param("id")

	//业务逻辑
	result := (&DemoLogic.SampleLogic{}).Read(context, id)

	respond.Json(context).Read(result)
}

//列表

func (instance *SampleController) Index(context *gin.Context) {
	//query
	queryStd := DemoQuery.SampleQuery{}
	query := request.Query(context).Get(&queryStd)

	//业务逻辑
	result, mata := (&DemoLogic.SampleLogic{}).Index(context, query)

	respond.Json(context).Index(result, mata)
}
