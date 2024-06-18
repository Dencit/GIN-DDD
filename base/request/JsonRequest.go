package request

import (
	BaseErr "app/base/err"
	"app/base/exception"
	"app/extend/convert/arrs"
	"app/extend/convert/structs"
	"github.com/gin-gonic/gin"
)

/**
notes: JSON输入类
*/

//JSON输入类-结构

type JsonRequestStruct struct {
	BaseRequestInterface
	Ctx *gin.Context
}

//JSON输入-实例

func Json(context *gin.Context) *JsonRequestStruct {
	instance := &JsonRequestStruct{}
	instance.Ctx = context
	return instance
}

//接收所有参数

func (instance *JsonRequestStruct) Input(std interface{}) map[string]any {
	err := instance.Ctx.ShouldBindJSON(&std)
	if err != nil {
		Code, _ := BaseErr.Root("VALIDATION_ERROR")
		Message := err.Error()
		exception.App(instance.Ctx, Code, Message)
	}
	jsonMap := structs.ToMap(std)
	return jsonMap
}

//只接收指定参数

func (instance *JsonRequestStruct) Only(std interface{}, Keys []string) map[string]any {
	err := instance.Ctx.ShouldBindJSON(&std)
	if err != nil {
		Code, _ := BaseErr.Root("VALIDATION_ERROR")
		Message := err.Error()
		exception.App(instance.Ctx, Code, Message)
	}
	jsonMap := structs.ToMap(std)
	for jsonKey := range jsonMap {
		match := arrs.InArray(jsonKey, Keys)
		if match == false {
			delete(jsonMap, jsonKey)
		}
	}
	return jsonMap
}

//不接收指定参数

func (instance *JsonRequestStruct) Except(std interface{}, Keys []string) map[string]any {
	err := instance.Ctx.ShouldBindJSON(&std)
	if err != nil {
		Code, _ := BaseErr.Root("VALIDATION_ERROR")
		Message := err.Error()
		exception.App(instance.Ctx, Code, Message)
	}
	jsonMap := structs.ToMap(std)
	for jsonKey := range jsonMap {
		match := arrs.InArray(jsonKey, Keys)
		if match == true {
			delete(jsonMap, jsonKey)
		}
	}
	return jsonMap
}
