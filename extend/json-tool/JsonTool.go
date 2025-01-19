package JsonTool

import (
	"app/extend/convert/maps"
	"encoding/json"
	"errors"
	"log"
)

//检查json文本格式

func IsJSONString(jsonStr string) bool {
	var err error
	var jsonMap json.RawMessage
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		return false
	}
	return true
}

//任意对象 转 json文本

func Encode(obj any) string {
	var err error
	var jsonBytes json.RawMessage
	jsonBytes, err = json.Marshal(obj)
	if err != nil {
		log.Println("JSON marshaling failed: %s", err)
	}
	return string(jsonBytes)
}

//json文本 转 字典

func MapDecode(jsonStr string) any {
	var err error
	var jsonMap json.RawMessage
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		log.Println("JSON unmarshaling failed: %s", err)
	}
	return jsonMap
}

//json文本 转 字典

func StructDecode(jsonStr string, std interface{}) any {
	var err error
	var jsonMap json.RawMessage
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		log.Println("JSON unmarshaling failed: %s", err)
	}

	err = maps.ToStruct(jsonMap, std)
	if err != nil {
		return errors.New("request data fail::" + err.Error())
	}

	return std
}
