package repo

import (
	BaseErr "app/base/err"
	"app/base/exception"
	BaseRepo "app/base/repo"
	UserEntity "app/domain/user/entity"
	"app/extend/convert/arrays"
	"app/extend/convert/maps"
	"app/extend/convert/values"
	MatchQuery "app/extend/match-query"
	"github.com/gin-gonic/gin"
	"math"
)

/**
notes: 领域层-仓储类
说明: 只写数据操作,不写别的内容,对应同名Entity 模型
调用原则: 向下调用[模型类]
*/

//gin结构

var Context *gin.Context

//仓储结构

type UserRepoStruct struct {
	BaseRepo.BaseRepoInterface
}

//仓储实例

func UserRepo(context *gin.Context) *UserRepoStruct {
	Context = context
	instance := &UserRepoStruct{&BaseRepo.BaseRepoStruct{}}
	return instance
}

//新增

func (receiver *UserRepoStruct) SaveOrFail(input map[string]any) interface{} {
	entity := UserEntity.User{
		Name: values.ToString(input["name"]),
	}
	builder := UserEntity.UserEntity()
	result := builder.Create(&entity)
	if result.Error != nil {
		Code, Message := BaseErr.Root("SAVE_FAIL")
		exception.App(Context, Code, Message)
		return nil
	}
	input["id"] = entity.ID
	return input
}

//更新

func (receiver *UserRepoStruct) UpdateOrFail(id string, input map[string]interface{}) interface{} {
	builder := UserEntity.UserEntity()
	builder.Where("id = ?", id)
	result := builder.Updates(&input)
	if result.Error != nil {
		Code, Message := BaseErr.Root("UPDATE_FAIL")
		exception.App(Context, Code, Message)
		return nil
	}
	input["id"] = id
	return input
}

//删除

func (receiver *UserRepoStruct) DeleteOrFail(id string) interface{} {
	input := make(map[string]any)
	builder := UserEntity.UserEntity()
	builder.Where("id = ?", id)
	result := builder.Delete(&input)
	if result.Error != nil {
		Code, Message := BaseErr.Root("DELETE_FAIL")
		exception.App(Context, Code, Message)
		return nil
	}
	input["id"] = id
	return input
}

//详情

func (receiver *UserRepoStruct) Read(id string) interface{} {
	entity := UserEntity.User{}
	builder := UserEntity.UserEntity()
	builder.Where("id = ?", id)
	builder.Order("updated_at Desc")
	result := builder.First(&entity)
	if result.Error != nil {
		//未找到数据
		return nil
	}
	return entity
}

//列表筛选

func (receiver *UserRepoStruct) Index(matchQuery *MatchQuery.MatchQueryStruct) (interface{}, map[string]any) {

	//实例化模型实体
	var entityList []UserEntity.User
	builder := UserEntity.UserEntity()

	//根据 ?_search=default 参数, 切换 捕捉到 ?type=1&status=1 ...的值的运算符.
	rule := make(map[string]string)
	action := matchQuery.SearchAction()
	if action == "default" {
		//rule["id"] = "="
		//rule["name"] = "like"
	}

	//捕捉 ?type=1&status=1 ... 的值, 转化成查询数组
	filterArr := make(map[string]string)
	searchArr := matchQuery.Search(rule, filterArr)
	if !maps.IsEmpty(searchArr) {
		maps.Walk(searchArr, func(value any, keyName any) {
			val, key := values.ToString(value), values.ToString(keyName)
			//自动添加查询条件
			builder.Where(key, val)
		})
	}

	//?_where_in_sort=status/1,2,3 //按id顺序返回结果
	// todo::待定

	//?_include=user,info - 副表关联模型,用于数据输出,不是查询条件.
	includeArr := matchQuery.Include()
	if !arrays.IsEmpty(includeArr) {
		arrays.Walk(includeArr, func(value any, index any) {
			val := values.ToString(value)
			builder.Preload(val)
		})
	}

	//?_sort = -id
	sortStr := matchQuery.Sort()
	if !values.IsEmpty(sortStr) {
		builder.Order(sortStr)
	}
	//默认排序
	builder.Order("updated_at Desc")

	//?_pagination = true
	meta, metaMap := matchQuery.Pagination()
	if meta.Pagination {
		builder.Offset(meta.Offset).Limit(meta.PageSize)
		builder.Count(&meta.Total)
		meta.PageTotal = math.Ceil(float64(meta.Total) / float64(meta.PageSize))
		metaMap["page_total"] = math.Ceil(float64(meta.Total) / float64(meta.PageSize))
		metaMap["total"] = meta.Total
	}

	//执行查询
	result := builder.Find(&entityList)
	if result.Error != nil {
		return nil, nil
	}
	return entityList, metaMap
}

//不存在则拦截

func (receiver *UserRepoStruct) IsNotExit(id string) interface{} {
	entity := UserEntity.User{}
	builder := UserEntity.UserEntity()
	builder.Where("id = ?", id)
	result := builder.First(&entity)
	if result.Error != nil {
		//未找到数据
		Code, Message := BaseErr.Root("ID_IS_NOT_EXIST")
		exception.App(Context, Code, Message)
		return nil
	}
	return entity
}

//是否存在

func (receiver *UserRepoStruct) IsHave(id string) interface{} {
	entity := UserEntity.User{}
	builder := UserEntity.UserEntity()
	builder.Where("id = ?", id)
	result := builder.First(&entity)
	if result.Error != nil {
		//未找到数据
		return nil
	}
	return entity
}
