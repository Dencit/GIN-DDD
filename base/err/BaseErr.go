package err

/**
notes: 错误码基础
*/

//声明结构

type BaseErrStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//业务码

func Root(Key string) (Code int, Message string) { // 创建字典实例
	var Err = make(map[string]BaseErrStruct)

	Err["VALIDATION_ERROR"] = BaseErrStruct{Code: 1026, Message: "输入参数有误"}
	//全局-公共业务码
	Err["AUTH_MUST"] = BaseErrStruct{Code: 1000000, Message: "访问必须授权"}
	Err["SAVE_FAIL"] = BaseErrStruct{Code: 1000011, Message: "数据新增失败"}
	Err["UPDATE_FAIL"] = BaseErrStruct{Code: 1000012, Message: "数据更新失败"}
	Err["DELETE_FAIL"] = BaseErrStruct{Code: 1000013, Message: "数据删除失败"}
	Err["ID_IS_NOT_EXIST"] = BaseErrStruct{Code: 1000014, Message: "ID 不存在"}
	Err["ID_IS_EXIST"] = BaseErrStruct{Code: 1000015, Message: "ID 已存在"}

	return Err[Key].Code, Err[Key].Message
}
