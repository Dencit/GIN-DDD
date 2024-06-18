package request

/**
notes: 应用层-输入验证类
desc: 只在此类 统一校验输入数据.
*/

//新增场景

type SampleSave struct {
	Name   string `form:"name" validate:"alphanum"`
	Mobile string `form:"mobile" validate:"alphanum"`
	Sex    uint8  `form:"sex" validate:"number"`
}

//更新场景

type SampleUpdate struct {
	Name   string `form:"name" validate:"omitempty,alphanum"`
	Mobile string `form:"mobile" validate:"omitempty,alphanum"`
	Sex    uint8  `form:"sex" validate:"omitempty,number"`
	Photo  string `validate:"omitempty,url"`
}
