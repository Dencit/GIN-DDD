package err

import BaseErr "app/base/err"

/**
notes: 根模块-总错误码
desc: 错误码区间,根据模块下的 doc.md 定义来设置. 注意 按数据单元做好注释, 每个单元错误码预留20位数间隔.
*/

//业务码

func Root(Key string) (Code int, Message string) { // 创建字典实例
	var Err = make(map[string]BaseErr.BaseErrStruct)

	//模块-公共业务码

	//单元-模型业务码
	Err["ID_IS_NOT_EXIST"] = BaseErr.BaseErrStruct{Code: 2000001, Message: "某某ID 不存在"}
	Err["ID_IS_EXIST"] = BaseErr.BaseErrStruct{Code: 2000002, Message: "某某ID 已存在"}

	return Err[Key].Code, Err[Key].Message
}
