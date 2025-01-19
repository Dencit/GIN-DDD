package respond

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
notes: http输出类基础
*/

//json输出类-结构

type HttpRespondStruct struct {
	BaseRespondInterface
	Ctx *gin.Context
}

//json输出类-实例

func Json(context *gin.Context) *HttpRespondStruct {
	instance := &HttpRespondStruct{}
	instance.Ctx = context
	return instance
}

//新增数据结果返回

func (receiver *HttpRespondStruct) Save(tableEntity interface{}) {
	receiver.Respond(receiver.Ctx, tableEntity, http.StatusCreated)
}

//更新数据结果返回

func (receiver *HttpRespondStruct) Update(tableEntity interface{}) {
	receiver.Respond(receiver.Ctx, tableEntity, http.StatusAccepted)
}

//删除数据结果返回

func (receiver *HttpRespondStruct) Delete() {
	receiver.Respond(receiver.Ctx, nil, http.StatusNoContent)
}

//输出单行数组

func (receiver *HttpRespondStruct) Read(tableEntity interface{}) {
	receiver.Respond(receiver.Ctx, tableEntity, http.StatusOK)
}

//输出多行数组

func (receiver *HttpRespondStruct) Index(result any, meta any) {
	receiver.RespondCollect(receiver.Ctx, result, meta, http.StatusOK)
}

//一般输出

func (receiver *HttpRespondStruct) Respond(context *gin.Context, tableEntity interface{}, httpStatus int) {
	//定义结构
	var respondJson = &respondStruct{}
	//默认值
	respondJson.Code = 0
	respondJson.Message = "success"
	respondJson.Data = nil
	respondJson.Meta = nil
	//数据非空
	if tableEntity != nil {
		respondJson.Data = tableEntity
	}
	context.JSON(httpStatus, respondJson)
	return
}

//多行输出+meta

func (receiver *HttpRespondStruct) RespondCollect(context *gin.Context, tableCollectStruct any, metas any, httpStatus int) {
	//定义结构
	var respondJson = &respondStruct{}
	//默认值
	respondJson.Code = 0
	respondJson.Message = "success"
	respondJson.Data = nil
	respondJson.Meta = nil
	//数据非空
	if tableCollectStruct != nil {
		respondJson.Data = tableCollectStruct
		respondJson.Meta = metas
	}
	context.JSON(httpStatus, respondJson)
	return
}
