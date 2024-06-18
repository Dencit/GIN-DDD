package query

/**
notes: 应用层-查询结构类
desc: 只在此类 统一查询结构.
*/

//筛选参数

type UserQuery struct {
	ID string `from:"id"`

	Role         string `from:"role"`
	Name         string `form:"name"`
	NickName     string `from:"nick_name"`
	Mobile       string `form:"mobile"`
	Avatar       string `from:"avatar"`
	PassWord     string `from:"pass_word"`
	ClientDriver string `from:"client_driver"`

	Sex        string `form:"sex"`
	Status     string `form:"status"`
	ClientType uint8  `from:"client_type"`

	Lat float32 `from:"lat"`
	Lng float32 `from:"lng"`

	OnLineTime  string `from:"on_line_time"`
	OffLineTime string `from:"off_line_time"`

	CreatedAt string `form:"created_at"`
	UpdatedAt string `form:"updated_at"`
	DeletedAt string `form:"deleted_at"`
}
