package maps

import "fmt"

// 检查 字典 为空
func IsEmpty[K comparable, V comparable](strMap map[K]V) bool {
	if strMap == nil || len(strMap) == 0 {
		return true
	}
	return false
}

func IsSet(string any) bool {
	str := fmt.Sprintf("%v", string)
	if string == nil || len(str) == 0 {
		return false
	}
	return true
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

// 闭包声明
type WalkCallback func(value any, key any)

// 闭包遍历数组
func Walk[K comparable, V comparable](array map[K]V, callback WalkCallback) {
	for key, value := range array {
		//log.Println("call::", value, key)//
		callback(value, key)
	}
}
