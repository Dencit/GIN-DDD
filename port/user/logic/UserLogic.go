package logic

import (
	UserRepo "app/domain/user/repo"
	MatchQuery "app/extend/match-query"
	"github.com/gin-gonic/gin"
	"log"
)

/**
notes: 应用层-业务类
说明: 业务类数据操作,一般不直接调用模型,通过仓储类提供存粹的数据执行函数, 跨 应用端/模块 操作同一数据类型的业务, 建议抽象到 领域层-业务类, 减少冗余.
调用原则: 向下调用[仓储类,领域层-业务类]
*/

//业务类结构

type UserLogic struct{}

//新增

func (receiver *UserLogic) Save(context *gin.Context) {

}

//更新

func (receiver *UserLogic) Update(context *gin.Context) {

}

//删除

func (receiver *UserLogic) Delete(context *gin.Context) {

}

//详情

func (receiver *UserLogic) Read(context *gin.Context, id string) interface{} {

	userRepo := UserRepo.UserRepo(context)
	result := userRepo.Read(id)

	return result
}

//列表

func (receiver *UserLogic) Index(context *gin.Context, query map[string]any) (interface{}, map[string]any) {

	//主表筛选逻辑-获取query查询表达式参数
	matchQuery := MatchQuery.Instance(query)
	log.Println("query::", query) //

	repo := UserRepo.UserRepo(context)
	result, meta := repo.Index(matchQuery)

	return result, meta
}
