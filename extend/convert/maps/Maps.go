package maps

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

/**
desc: 字典转换
*/

//检查字典为空

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

//获取字典键值

func Keys(mapArr map[string]any) sort.StringSlice {
	var keys sort.StringSlice
	for key := range mapArr {
		keys = append(keys, key)
	}
	return keys
}

func DiffKey(mapArr1 map[string]any, mapArr2 map[string]any) map[string]any {
	diffMap := make(map[string]any)
	for keyName, value := range mapArr1 {
		diffMap[keyName] = value
		//右边存在,则删除.
		if InMap(keyName, mapArr2) == true {
			delete(diffMap, keyName)
		}
	}
	return diffMap
}

func InMap(keyName string, strMap map[string]any) bool {
	res := false
	for key, _ := range strMap {
		if keyName == key {
			res = true
			break
		}
	}
	return res
}

//闭包声明

type walkCallback func(value any, keyName any)

//闭包遍历数组

func Walk[K comparable, V comparable](array map[K]V, callback walkCallback) {
	for key, value := range array {
		//log.Println("call::", value, key)//
		callback(value, key)
	}
}

//字典升序 - 通过回调方法 取参数

func KSort(mapArr map[string]any, callback walkCallback) {
	//获取字典键值
	keys := Keys(mapArr)
	//升序
	sort.Sort(keys)
	//赋值
	for _, keyName := range keys {
		callback(mapArr[keyName], keyName)
	}
	return
}

//字典倒序 - 通过回调方法 取参数

func RSort(mapArr map[string]any, callback walkCallback) {
	//获取字典键值
	keys := Keys(mapArr)
	//升序
	sort.Sort(sort.Reverse(keys))
	//赋值
	for _, keyName := range keys {
		callback(mapArr[keyName], keyName)
	}
	return
}

//字典转结构

func ToStruct(ParamsMap interface{}, std interface{}) error {
	jsonStr, err := json.Marshal(ParamsMap)
	if err != nil {
		return errors.New("json marshal fail")
	}
	err = json.Unmarshal(jsonStr, std)
	if err != nil {
		return errors.New("json unmarshal fail " + err.Error())
	}
	return nil
}
