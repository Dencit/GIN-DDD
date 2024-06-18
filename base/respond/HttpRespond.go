package respond

import (
	"app/extend/convert/ints"
	"encoding/json"
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
	jsonBytes, _ := json.Marshal(tableEntity)
	jsonStr := string(jsonBytes)
	json.Unmarshal([]byte(jsonStr), &tableEntity)
	receiver.Respond(respCtx, tableEntity, http.StatusCreated)
}

//更新数据结果返回

func (receiver *HttpRespondStruct) Update(tableEntity interface{}) {
	jsonBytes, _ := json.Marshal(tableEntity)
	jsonStr := string(jsonBytes)
	json.Unmarshal([]byte(jsonStr), &tableEntity)
	receiver.Respond(respCtx, tableEntity, http.StatusAccepted)
}

//输出数据结果返回

func (receiver *HttpRespondStruct) Delete() {
	receiver.Respond(respCtx, nil, http.StatusNoContent)
}

//输出单行数组

func (receiver *HttpRespondStruct) Read(tableEntity interface{}) {

	jsonBytes, _ := json.Marshal(tableEntity)
	jsonStr := string(jsonBytes)
	json.Unmarshal([]byte(jsonStr), &tableEntity)
	receiver.Respond(respCtx, tableEntity, http.StatusOK)
}

//输出多行数组

func (receiver *HttpRespondStruct) Index(tableCollectEntity interface{}, metaMap map[string]any) {

	fieldJsonBytes, _ := json.Marshal(tableCollectEntity)
	fieldJsonStr := string(fieldJsonBytes)
	json.Unmarshal([]byte(fieldJsonStr), &tableCollectEntity)

	receiver.RespondCollect(respCtx, tableCollectEntity, metaMap, http.StatusOK)
}

//输出多行数组+翻页字段

func (receiver *HttpRespondStruct) IndexPage(fields interface{}, tableCollectStruct interface{}) {

	fieldJsonBytes, _ := json.Marshal(fields)
	fieldJsonStr := string(fieldJsonBytes)
	json.Unmarshal([]byte(fieldJsonStr), &tableCollectStruct)

	metaJsonStr := `{ "pagination":true , "per_page":20, "page":1 }`
	var metaMap map[string]interface{}
	json.Unmarshal([]byte(metaJsonStr), &metaMap)

	perPage := ints.ToIntDefault(respCtx.Query("_per_page"), 20)
	page := ints.ToIntDefault(respCtx.Query("_page"), 1)

	metaMap["pagination"] = true
	metaMap["per_page"] = perPage
	metaMap["page"] = page

	receiver.RespondCollect(respCtx, tableCollectStruct, metaMap, http.StatusOK)
}
