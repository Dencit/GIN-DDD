package query

/**
notes: 应用层-查询结构类
desc: 只在此类 统一查询结构.
*/

//筛选参数

type SampleQuery struct {
	ID string `from:"id"`

	Name   string `form:"name"`
	Mobile string `form:"mobile"`
	Photo  string `form:"photo"`
	Sex    string `form:"sex"`
	Type   string `form:"type"`
	Status string `form:"status"`

	CreatedAt string `form:"created_at"`
	UpdatedAt string `form:"updated_at"`
	DeletedAt string `form:"deleted_at"`
}
