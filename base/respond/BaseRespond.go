package respond

import (
	"github.com/gin-gonic/gin"
)

/**
notes: 输出类基础
*/

//输出类-接口

type BaseRespondInterface interface {
	Respond(context *gin.Context, tableEntity interface{}, httpStatus int)
	RespondCollect(context *gin.Context, tableCollectStruct interface{}, metas map[string]any, httpStatus int)
}

//输出类-结构

type BaseRespondStruct struct {
	BaseRespondInterface
}

//一般输出

func (instance *BaseRespondStruct) Respond(context *gin.Context, tableEntity interface{}, httpStatus int) {
	//定义结构
	var jsonStruct struct {
		Code    uint        `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Mata    interface{} `json:"mata"`
	}
	//默认值
	jsonStruct.Code = 0
	jsonStruct.Message = "success"
	jsonStruct.Data = nil
	jsonStruct.Mata = nil
	//数据非空
	if tableEntity != nil {
		jsonStruct.Data = tableEntity
	}
	context.JSON(httpStatus, jsonStruct)
	return
}

//多行输出+meta

func (instance *BaseRespondStruct) RespondCollect(context *gin.Context, tableCollectStruct interface{}, metas map[string]any, httpStatus int) {
	//定义结构
	var jsonStruct struct {
		Code    uint        `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Mata    interface{} `json:"mata"`
	}
	//默认值
	jsonStruct.Code = 0
	jsonStruct.Message = "success"
	jsonStruct.Data = nil
	jsonStruct.Mata = nil
	//数据非空
	if tableCollectStruct != nil {
		jsonStruct.Data = tableCollectStruct
		jsonStruct.Mata = metas
	}
	context.JSON(httpStatus, jsonStruct)
	return
}
