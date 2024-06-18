package match_query

import (
	"app/extend/convert/arrs"
	"app/extend/convert/maps"
	"app/extend/convert/strs"
	"app/extend/convert/structs"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

type QueryMatchStructInterface interface {
}

type MatchQueryStruct struct {
	QueryMatchStructInterface
	RequestQuery map[string]any
}

type MataStruct struct {
	Pagination bool `json:"pagination"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	Offset     int  `json:"offset"`
	//PageTotal  int  `json:"page_total"`
	Total int64 `json:"total"`
}

func Instance(requestQuery map[string]any) *MatchQueryStruct {
	instance := &MatchQueryStruct{}
	instance.RequestQuery = requestQuery
	return instance
}

func (instance *MatchQueryStruct) SearchAction() string {
	var action string = "default"
	if instance.RequestQuery["_search"] != "" {
		action = strs.ToStr(instance.RequestQuery["_search"])
	}
	return action
}

func (instance *MatchQueryStruct) Search(rule map[string]string, filterArr map[string]string) map[string]any {
	//返回结果
	searchArr := make(map[string]any, 0)

	//排除参数集
	outQuery := []any{
		"_pagination", "_page", "_page_size",
		"_where", "_where_in", "where_in_sort", "_include", "_extend", "_search",
		"_sort", "_group", "_time",
	}

	flipMap := arrs.Flip(outQuery)
	queryArr := maps.DiffKey(instance.RequestQuery, flipMap)

	if !maps.IsEmpty(filterArr) {
		// todo::待定
	}
	if !maps.IsEmpty(queryArr) {
		maps.Walk(queryArr, func(value any, key any) {
			//log.Println("back::", value, key)//
			keyName := strs.ToStr(key)
			valueName := strs.ToStr(value)
			operator := "="
			if maps.IsSet(rule[keyName]) {
				operator = rule[keyName]
			}
			//log.Println("operator::", operator)//
			//筛选运算符预处理
			ope, val := instance.searchOperator(operator, valueName)
			//值非空字符串才获取
			if valueName != "" {
				//currArr := []any{keyName + " " + ope + " ?", val}
				//searchArr = append(searchArr, currArr)
				currKey := keyName + " " + ope + " ?"
				searchArr[currKey] = val
			}
		})
		//log.Println("queryArr::", queryArr)//

	}
	return searchArr
}

func (instance *MatchQueryStruct) Include() []string {
	var includeArr []string
	if !strs.IsEmpty(instance.RequestQuery["_include"]) {
		joinStr := strs.ToStr(instance.RequestQuery["_include"])
		joinArr := strings.Split(joinStr, ",")
		//首字母大写
		loader := cases.Title(language.Und, cases.NoLower)
		for index, value := range joinArr {
			joinArr[index] = loader.String(value)
		}
		includeArr = joinArr
	}
	return includeArr
}

func (instance *MatchQueryStruct) searchOperator(operator string, value any) (string, any) {
	val := fmt.Sprintf("%v", value)
	switch operator {
	case "like": //模糊筛选处理
		value = "%" + val + "%"
		matchL, _ := regexp.MatchString("^(%|\\*)", val)
		if matchL {
			value = "%" + val
		}
		matchR, _ := regexp.MatchString("(%|\\*)$", val)
		if matchR {
			value = val + "%"
		}
		break
	case "=": //兼容多选-逗号分隔
		regex := regexp.MustCompile("(,|\\%|\\*)")
		match := regex.FindString(val)
		if match == "," {
			operator = "IN"
			split := strings.Split(val, ",")
			value = split
		}
		if match == "%" {
			operator = "LIKE"
		}
		if match == "*" {
			operator = "LIKE"
			value = regex.ReplaceAllString(val, "%")
		}
		break
	}

	return operator, value
}

func (instance *MatchQueryStruct) Sort() string {
	sortStr := ""
	if !strs.IsEmpty(instance.RequestQuery["_sort"]) {
		orderStr := strs.ToStr(instance.RequestQuery["_sort"])
		sortStr = instance.sortOperator(orderStr)
	}
	return sortStr
}

// 排序-sort参数转换
func (instance *MatchQueryStruct) sortOperator(orderStr string) string {
	sortStr := ""
	orderArr := arrs.Explode(",", orderStr)
	orderType := "ASC"
	for _, value := range orderArr {
		regex := regexp.MustCompile("^(-|)")
		match := regex.FindString(value)
		if match == "-" {
			orderType = "DESC"
			value = strings.TrimLeft(value, "-")
		}
		sortStr += value + " " + orderType + ", "
	}
	sortStr = strings.TrimRight(sortStr, ", ")
	return sortStr
}

func (instance *MatchQueryStruct) Pagination() (mataStruct MataStruct, metaMap map[string]any) {
	//默认值
	pagination := true
	page, pageSize := 1, 20
	offset := 100
	//获取参数
	if !strs.IsEmpty(instance.RequestQuery["_pagination"]) {
		pagination = strs.ToBool(instance.RequestQuery["_pagination"])
	}
	if !strs.IsEmpty(instance.RequestQuery["_page"]) {
		page = strs.ToInt(instance.RequestQuery["_page"])
	}
	if !strs.IsEmpty(instance.RequestQuery["_page_size"]) {
		pageSize = strs.ToInt(instance.RequestQuery["_page_size"])
	}
	//限制范围
	if page < 1 {
		page = 1
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset = (page - 1) * pageSize

	mataStruct = MataStruct{
		Pagination: pagination,
		Page:       page,
		PageSize:   pageSize,
		Offset:     offset,
		//PageTotal:  0,
		Total: 0,
	}
	metaMap = structs.ToMap(&mataStruct)

	return mataStruct, metaMap
}
