package respond

import (
	"github.com/gin-gonic/gin"
)

/**
notes: 输出类基础
*/

//定义返回结构

type respondStruct struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

//输出类-接口

type BaseRespondInterface interface {
	Respond(context *gin.Context, tableEntity interface{}, httpStatus int)
	RespondCollect(context *gin.Context, tableCollectStruct any, metas any, httpStatus int)
}

//输出类-结构

type BaseRespondStruct struct {
	BaseRespondInterface
}

//一般输出

func (receiver *BaseRespondStruct) Respond(context *gin.Context, tableEntity interface{}, httpStatus int) {
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

func (receiver *BaseRespondStruct) RespondCollect(context *gin.Context, tableCollectStruct any, metas any, httpStatus int) {
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
