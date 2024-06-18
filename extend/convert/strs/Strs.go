package strs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsEmpty(string any) bool {
	str := fmt.Sprintf("%v", string)
	if string == nil || len(str) == 0 {
		return true
	}
	return false
}

func ToBool(string any) bool {
	str := fmt.Sprintf("%v", string)
	boo, _ := strconv.ParseBool(str)
	return boo
}

func ToInt(string any) int {
	str := fmt.Sprintf("%v", string)
	number, _ := strconv.Atoi(str)
	return number
}

func ToStr(string any) string {
	str := fmt.Sprintf("%v", string)
	return str
}

func ToStrDefault(string interface{}, defaultStr string) string {
	str := fmt.Sprintf("%v", string)
	if str == "" {
		return defaultStr
	}
	return str
}

// 字符串 转 小写下划线
func toSnakeCase(str string, private bool) string {
	var matchFirstCap = regexp.MustCompile("([A-Z])")
	str = matchFirstCap.ReplaceAllString(str, "_${1}")
	str = strings.ToLower(str)
	if private == false {
		str = strings.TrimLeft(str, "_")
	}
	return str
}
