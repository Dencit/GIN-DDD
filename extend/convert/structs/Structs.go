package structs

import (
	"fmt"
	"reflect"
)

/**
desc: 结构转换
*/

//检查结构为空

func IsEmpty[K comparable, V comparable](strMap map[K]V) bool {
	if strMap == nil || len(strMap) == 0 {
		return true
	}
	return false
}

//检查存在值

func IsSet(string any) bool {
	str := fmt.Sprintf("%v", string)
	if string == nil || len(str) == 0 {
		return false
	}
	return true
}

//结构转字典

func ToMap(content interface{}) map[string]any {
	var resMap = make(map[string]any)

	if content != nil {
		//通过反射获取 type,value 定义
		typeElem := reflect.TypeOf(content).Elem()
		valueElem := reflect.ValueOf(content).Elem()
		for i := 0; i < typeElem.NumField(); i++ {
			currType := typeElem.Field(i)
			fromTag := currType.Tag.Get("form") //from tag
			jsonTag := currType.Tag.Get("json") //json tag
			currValue := valueElem.Field(i)
			value := currValue.Interface()
			if value != "" {
				if fromTag != "" {
					resMap[fromTag] = value
				} else if jsonTag != "" {
					resMap[jsonTag] = value
				}
			}
		}
	}

	return resMap
}
