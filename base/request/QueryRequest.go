package request

import (
	"app/extend/convert/arrs"
	"app/extend/convert/structs"
	"github.com/gin-gonic/gin"
	"log"
)

/**
notes: 路径参数输入类
*/

//默认路径参数

type DefQuery struct {
	Search  string `form:"_search"`  //筛选条件场景
	Include string `form:"_include"` //关联查询数据
	Extend  string `form:"_extend"`  //扩展查询条件
	Sort    string `form:"_sort"`    //排序
	Group   string `form:"_group"`   //分组
	//
	Pagination string `form:"_pagination"` //翻页开关
	Page       string `form:"_page"`       //当前页
	PageSize   string `form:"_page_size"`  //每页数量
	//
	Time string `form:"_time"` //缓存开关
}

//路径参数输入类-结构

type QueryRequestStruct struct {
	BaseRequestInterface
	Ctx      *gin.Context
	QueryMap map[string]any
}

//路径参数输入-实例

func Query(context *gin.Context) *QueryRequestStruct {
	instance := &QueryRequestStruct{}
	instance.Ctx = context

	//暂存默认参数
	instance.QueryMap = nil
	QueryStruct := DefQuery{}
	err := instance.Ctx.ShouldBindQuery(&QueryStruct)
	if err != nil {
		log.Println("Query::", err.Error())
	}
	instance.QueryMap = structs.ToMap(&QueryStruct)

	return instance
}

//获取所有参数

func (instance *QueryRequestStruct) Get(std interface{}) map[string]any {
	err := instance.Ctx.ShouldBindQuery(std)
	if err != nil {
		log.Println("SetQuery::", err.Error())
	}
	valueQueryMap := structs.ToMap(std)
	instance.QueryMap = arrs.Merge(valueQueryMap, instance.QueryMap)
	return instance.QueryMap
}

//只获取指定参数

func (instance *QueryRequestStruct) Only(std interface{}, Keys []string) map[string]any {
	err := instance.Ctx.ShouldBindQuery(std)
	if err != nil {
		log.Println("SetQuery::", err.Error())
	}
	valueQueryMap := structs.ToMap(std)
	instance.QueryMap = arrs.Merge(valueQueryMap, instance.QueryMap)
	for QueryKey := range instance.QueryMap {
		match := arrs.InArray(QueryKey, Keys)
		if match == false {
			delete(instance.QueryMap, QueryKey)
		}
	}
	return instance.QueryMap
}

//不获取指定参数

func (instance *QueryRequestStruct) Except(std interface{}, Keys []string) map[string]any {
	err := instance.Ctx.ShouldBindQuery(std)
	if err != nil {
		log.Println("SetQuery::", err.Error())
	}
	valueQueryMap := structs.ToMap(std)
	instance.QueryMap = arrs.Merge(valueQueryMap, instance.QueryMap)
	for QueryKey := range instance.QueryMap {
		match := arrs.InArray(QueryKey, Keys)
		if match == true {
			delete(instance.QueryMap, QueryKey)
		}
	}
	return instance.QueryMap
}
