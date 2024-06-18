package structs

import (
	"app/extend/convert/strs"
	"reflect"
)

// 结构转字典
func ToMap(content interface{}) map[string]any {
	var resMap = make(map[string]any)
	//通过反射获取 type,value 定义
	types := reflect.TypeOf(content).Elem()
	values := reflect.ValueOf(content).Elem()
	for i := 0; i < types.NumField(); i++ {
		fromTag := types.Field(i).Tag.Get("form") //from tag
		jsonTag := types.Field(i).Tag.Get("json") //json tag
		value := values.Field(i).Interface()
		if value != "" {
			if fromTag != "" {
				resMap[fromTag] = strs.ToStr(value)
			} else if jsonTag != "" {
				resMap[jsonTag] = strs.ToStr(value)
			}
		}
	}
	return resMap
}
