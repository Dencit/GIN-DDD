package controller

import (
	"app/base/request"
	"app/base/respond"
	"app/base/validate"
	ApiCache "app/extend/api-cache"
	DemoLogic "app/port/demo/logic"
	DemoQuery "app/port/demo/query"
	DemoRequest "app/port/demo/request"
	"github.com/gin-gonic/gin"
	"time"
)

/**
notes: 应用层-控制器
说明: 控制器内不写业务,只写http层面相关的逻辑,
调用原则: 向下调用[输入验证类,业务类,输出转化类].
*/

//控制器结构

type SampleController struct{}

//新增

func (receiver *SampleController) Save(context *gin.Context) {

	//输入验证
	inputStd := DemoRequest.SampleSave{}
	input := request.Json(context).Input(&inputStd)
	validate.Check(context).Command(&inputStd)

	//业务逻辑
	result := (&DemoLogic.SampleLogic{}).Save(context, input)

	respond.Json(context).Save(result)
}

//更新

func (receiver *SampleController) Update(context *gin.Context) {

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

func (receiver *SampleController) Delete(context *gin.Context) {
	//路径ID
	id := context.Param("id")

	//业务逻辑
	(&DemoLogic.SampleLogic{}).Delete(context, id)

	respond.Json(context).Delete()
}

//详情

func (receiver *SampleController) Read(context *gin.Context) {
	//path
	id := context.Param("id")

	//query
	queryStd := DemoQuery.SampleQuery{}
	query := request.Query(context).Get(&queryStd)

	//检查url参数缓存
	apiCache := ApiCache.Instance(query)
	hKey := apiCache.HKeyByClassMethod("sample@read")
	queryKey := apiCache.QueryKeyByRequest(query) + "&id=" + id
	//缓存闭包
	result, _ := apiCache.Collect(hKey, queryKey, func(result *any, meta *any) {

		//业务逻辑 - 修改实参
		*result = (&DemoLogic.SampleLogic{}).Read(context, id)

	}, 300*time.Second)

	respond.Json(context).Read(result)
}

//列表

func (receiver *SampleController) Index(context *gin.Context) {
	//query
	queryStd := DemoQuery.SampleQuery{}
	query := request.Query(context).Get(&queryStd)

	//检查url参数缓存
	apiCache := ApiCache.Instance(query)
	hKey := apiCache.HKeyByClassMethod("sample@index")
	queryKey := apiCache.QueryKeyByRequest(query)
	//缓存闭包
	result, meta := apiCache.Collect(hKey, queryKey, func(result *any, meta *any) {

		//业务逻辑 - 修改实参
		*result, *meta = (&DemoLogic.SampleLogic{}).Index(context, query)

	}, 300*time.Second)

	respond.Json(context).Index(result, meta)
}
