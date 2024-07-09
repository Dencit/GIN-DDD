package structs

import (
	"app/extend/convert/values"
	"reflect"
)

/**
desc: 结构转换
*/

//结构转字典

func ToMap(content interface{}) map[string]any {
	var resMap = make(map[string]any)
	//通过反射获取 type,value 定义
	typeElem := reflect.TypeOf(content).Elem()
	valueElem := reflect.ValueOf(content).Elem()
	for i := 0; i < typeElem.NumField(); i++ {
		fromTag := typeElem.Field(i).Tag.Get("form") //from tag
		jsonTag := typeElem.Field(i).Tag.Get("json") //json tag
		value := valueElem.Field(i).Interface()
		if value != "" {
			if fromTag != "" {
				resMap[fromTag] = values.ToString(value)
			} else if jsonTag != "" {
				resMap[jsonTag] = values.ToString(value)
			}
		}
	}
	return resMap
}
