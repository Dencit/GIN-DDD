package arrays

import (
	"app/extend/convert/values"
	"fmt"
	"strings"
)

/**
desc: 切片数组转换
*/

//检查 切片数组 为空

func IsEmpty[T any](slice []T) bool {
	if slice == nil || len(slice) == 0 {
		return true
	}
	return false
}

//检查 切片数组 存在某值

func IsSet(string any) bool {
	str := fmt.Sprintf("%v", string)
	if string == nil || len(str) == 0 {
		return false
	}
	return true
}

//将字符串 按照指定的分隔符 拆分成一个数组

func Explode(sep string, orderStr any) []string {
	str := fmt.Sprintf("%v", orderStr)
	orderArr := strings.Split(str, sep)
	return orderArr
}

//切片数组 - 键值反转为 字典

func Flip(array []any) map[string]any {
	mapArr := make(map[string]any)
	for index, value := range array {
		val := values.ToString(value)
		ind := values.ToString(index)
		mapArr[val] = ind
	}
	return mapArr
}

func InArray(target string, strArray []string) bool {
	res := false
	for _, element := range strArray {
		if target == element {
			res = true
			break
		}
	}
	return res
}

//合并字典

func Merge(strMap1 map[string]any, strMap2 map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range strMap1 {
		result[k] = v
	}
	for k, v := range strMap2 {
		result[k] = v
	}
	return result
}

//闭包声明

type walkCallback func(value any, index any)

//闭包遍历数组

func Walk[T any](array []T, callback walkCallback) {
	for index, value := range array {
		//log.Println("call::", value, index) //
		callback(value, index)
	}
}
