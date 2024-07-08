package json_tool

import (
	"encoding/json"
	"log"
)

//任意对象 转 json文本

func Encode(obj any) string {
	var err error
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(obj)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	return string(jsonBytes)
}

//json文本 转 字典

func Decode(jsonStr string) any {
	var err error
	var jsonMap any
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	return jsonMap
}
