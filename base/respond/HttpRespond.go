package respond

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
notes: http输出类基础
*/

//gin结构

var respCtx *gin.Context

//json输出类-结构

type HttpRespondStruct struct {
	BaseRespondInterface
}

//json输出类-实例

func Json(context *gin.Context) *HttpRespondStruct {
	instance := &HttpRespondStruct{&BaseRespondStruct{}}
	respCtx = context
	return instance
}

//新增数据结果返回

func (receiver *HttpRespondStruct) Save(tableEntity interface{}) {
	receiver.Respond(respCtx, tableEntity, http.StatusCreated)
}

//更新数据结果返回

func (receiver *HttpRespondStruct) Update(tableEntity interface{}) {
	receiver.Respond(respCtx, tableEntity, http.StatusAccepted)
}

//输出数据结果返回

func (receiver *HttpRespondStruct) Delete() {
	receiver.Respond(respCtx, nil, http.StatusNoContent)
}

//输出单行数组

func (receiver *HttpRespondStruct) Read(tableEntity interface{}) {
	receiver.Respond(respCtx, tableEntity, http.StatusOK)
}

//输出多行数组

func (receiver *HttpRespondStruct) Index(result any, meta any) {
	receiver.RespondCollect(respCtx, result, meta, http.StatusOK)
}
