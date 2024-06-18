package validate

import (
	BaseErr "app/base/err"
	"app/base/exception"
	"app/extend/convert/strs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/**
notes: 输入验证类基础
*/

//输入验证类-接口

type BaseValidateInterface interface {
}

//输入验证类-结构

type BaseValidateStruct struct {
	BaseValidateInterface
	ValCtx *gin.Context
}

//输入验证类-实例

func Check(context *gin.Context) *BaseValidateStruct {
	instance := &BaseValidateStruct{}
	instance.ValCtx = context
	return instance
}

//加载验证设置

func (instance *BaseValidateStruct) Command(std interface{}) interface{} {
	validate := validator.New()
	err := validate.Struct(std)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			Code, _ := BaseErr.Root("VALIDATION_ERROR")
			Message := strs.ToStr(errors[0])
			exception.App(instance.ValCtx, Code, Message)
		}
	}
	return std
}
