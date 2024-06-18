package repo

import (
	BaseErr "app/base/err"
	"app/base/exception"
	BaseRepo "app/base/repo"
	DemoEntity "app/domain/demo/entity"
	"app/extend/convert/arrs"
	"app/extend/convert/maps"
	"app/extend/convert/strs"
	"app/extend/match-query"
	"github.com/gin-gonic/gin"
	"log"
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

type SampleRepoStruct struct {
	BaseRepo.BaseRepoInterface
}

//仓储实例

func SampleRepo(context *gin.Context) *SampleRepoStruct {
	Context = context
	instance := &SampleRepoStruct{&BaseRepo.BaseRepoStruct{}}
	return instance
}

//新增

func (receiver *SampleRepoStruct) SaveOrFail(input map[string]any) interface{} {
	entity := DemoEntity.Sample{
		Name: strs.ToStr(input["name"]),
	}
	builder := DemoEntity.SampleEntity()
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

func (receiver *SampleRepoStruct) UpdateOrFail(id string, input map[string]any) interface{} {
	builder := DemoEntity.SampleEntity()
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

func (receiver *SampleRepoStruct) DeleteOrFail(id string) interface{} {
	input := make(map[string]any)
	builder := DemoEntity.SampleEntity()
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

func (receiver *SampleRepoStruct) Read(id string) interface{} {
	entity := DemoEntity.Sample{}
	builder := DemoEntity.SampleEntity()
	builder.Where("id = ?", id)
	builder.Order("updated_at Desc")

	//ent := UserEntity.User{}      //
	//bu := UserEntity.UserEntity() //
	//bu.Where("id = ?", id)        //
	//res := bu.First(&ent)         //
	//log.Println("res::", res)     //

	result := builder.First(&entity)
	if result.Error != nil {
		//未找到数据
		return nil
	}
	return entity
}

//列表筛选

func (receiver *SampleRepoStruct) Index(matchQuery *match_query.MatchQueryStruct) (interface{}, map[string]any) {

	//实例化模型实体
	var entityList []DemoEntity.Sample
	builder := DemoEntity.SampleEntity()

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
	log.Println("searchArr::", searchArr) //
	if !maps.IsEmpty(searchArr) {
		maps.Walk(searchArr, func(value any, key any) {
			val, key := strs.ToStr(value), strs.ToStr(key)
			//自动添加查询条件
			builder.Where(key, val)
		})
	}

	//?_where_in_sort=status/1,2,3 //按id顺序返回结果
	// todo::待定

	//?_include=user,info - 副表关联模型,用于数据输出,不是查询条件.
	includeArr := matchQuery.Include()
	log.Println("includeArr::", includeArr) //
	if !arrs.IsEmpty(includeArr) {
		arrs.Walk(includeArr, func(value any, index any) {
			val := strs.ToStr(value)
			builder.Preload(val)
		})
	}

	//?_sort = -id
	sortStr := matchQuery.Sort()
	if !strs.IsEmpty(sortStr) {
		builder.Order(sortStr)
	}
	//默认排序
	builder.Order("updated_at Desc")

	//?_pagination = true
	mata, metaMap := matchQuery.Pagination()
	if mata.Pagination {
		builder.Offset(mata.Offset).Limit(mata.PageSize)
		builder.Count(&mata.Total)
		metaMap["page_total"] = math.Ceil(float64(mata.Total) / float64(mata.PageSize))
		metaMap["total"] = mata.Total
		log.Println("mataStruct", mata) //
		log.Println("metaMap", metaMap) //
	}

	//执行查询
	result := builder.Find(&entityList)
	if result.Error != nil {
		return nil, nil
	}
	return entityList, metaMap
}

//不存在则拦截

func (receiver *SampleRepoStruct) IsNotExit(id string) interface{} {
	entity := DemoEntity.Sample{}
	builder := DemoEntity.SampleEntity()
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

func (receiver *SampleRepoStruct) IsHave(id string) interface{} {
	entity := DemoEntity.Sample{}
	builder := DemoEntity.SampleEntity()
	builder.Where("id = ?", id)
	result := builder.First(&entity)
	if result.Error != nil {
		//未找到数据
		return nil
	}
	return entity
}
