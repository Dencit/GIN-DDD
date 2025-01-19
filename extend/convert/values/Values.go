package values

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
desc: 值转换,包括整型和字符串.
*/

//任意值 转 整型

func ToInt(obj any) int {
	objStr := fmt.Sprintf("%v", obj)
	_int, err := strconv.Atoi(objStr)
	if err != nil {
		return 0
	}
	return _int
}

//任意值 转 64位整型

func ToInt64(obj any) int64 {
	objStr := fmt.Sprintf("%v", obj)
	_int64, err := strconv.ParseInt(objStr, 10, 64)
	if err != nil {
		return 0
	}
	return _int64
}

//任意值 转 64位 无符号整型

func ToUint64(obj any) uint64 {
	objStr := fmt.Sprintf("%v", obj)
	_uint64, err := strconv.ParseUint(objStr, 10, 64)
	if err != nil {
		return 0
	}
	return _uint64
}

//任意值 转 字符串

func ToString(obj any) string {
	objStr := fmt.Sprintf("%v", obj)
	return objStr
}

//任意值 转 布尔值

func ToBool(obj any) bool {
	objStr := fmt.Sprintf("%v", obj)
	boo, _ := strconv.ParseBool(objStr)
	return boo
}

//检查字符串为空

func IsEmpty(string any) bool {
	str := fmt.Sprintf("%v", string)
	if string == nil || len(str) == 0 {
		return true
	}
	return false
}

//字符串 转 小写下划线

func toSnakeCase(str string, private bool) string {
	var matchFirstCap = regexp.MustCompile("([A-Z])")
	str = matchFirstCap.ReplaceAllString(str, "_${1}")
	str = strings.ToLower(str)
	if private == false {
		str = strings.TrimLeft(str, "_")
	}
	return str
}
