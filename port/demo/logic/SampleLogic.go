package logic

import (
	DemoRepo "app/domain/demo/repo"
	match_query "app/extend/match-query"
	"github.com/gin-gonic/gin"
	"log"
)

/**
notes: 应用层-业务类
说明: 业务类数据操作,一般不直接调用模型,通过仓储类提供存粹的数据执行函数, 跨 应用端/模块 操作同一数据类型的业务, 建议抽象到 领域层-业务类, 减少冗余.
调用原则: 向下调用[仓储类,领域层-业务类]
*/

//业务类结构

type SampleLogic struct{}

//新增

func (receiver *SampleLogic) Save(context *gin.Context, input map[string]any) interface{} {

	sampleRepo := DemoRepo.SampleRepo(context)
	result := sampleRepo.SaveOrFail(input)

	return result
}

//更新

func (receiver *SampleLogic) Update(context *gin.Context, id string, input map[string]any) interface{} {

	sampleRepo := DemoRepo.SampleRepo(context)
	sampleRepo.IsNotExit(id)

	result := sampleRepo.UpdateOrFail(id, input)

	return result
}

//删除

func (receiver *SampleLogic) Delete(context *gin.Context, id string) interface{} {

	sampleRepo := DemoRepo.SampleRepo(context)
	sampleRepo.IsNotExit(id)

	sampleRepo.DeleteOrFail(id)

	return nil
}

//详情

func (receiver *SampleLogic) Read(context *gin.Context, id string) interface{} {

	sampleRepo := DemoRepo.SampleRepo(context)
	result := sampleRepo.Read(id)

	return result
}

//列表

func (receiver *SampleLogic) Index(context *gin.Context, query map[string]any) (interface{}, map[string]any) {
	//主表筛选逻辑-获取query查询表达式参数
	matchQuery := match_query.Instance(query)
	log.Println("query::", query) //

	repo := DemoRepo.SampleRepo(context)
	result, mata := repo.Index(matchQuery)

	return result, mata
}
