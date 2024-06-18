package ints

import "strconv"

func Int64ToInt(i64 int64) int {
	strInt64 := strconv.FormatInt(i64, 10)
	_int, err := strconv.Atoi(strInt64)
	if err != nil {
		return 0
	}
	return _int
}

func ToInt(string string) int {
	_int, err := strconv.Atoi(string)
	if err != nil {
		return 0
	}
	return _int
}

func ToIntDefault(string string, defaultInt int) int {
	_int, err := strconv.Atoi(string)
	if err != nil {
		return defaultInt
	}
	return _int
}

func ToInt64(string string) int64 {
	_int64, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return 0
	}
	return _int64
}

func ToInt64Default(string string, defaultInt int64) int64 {
	_int64, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return defaultInt
	}
	return _int64
}
