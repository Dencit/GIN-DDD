package respond

/**
notes: 输出类基础
*/

//定义返回结构

type respondStruct struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

//输出类-接口

type BaseRespondInterface interface {
}

//输出类-结构

type BaseRespondStruct struct {
	BaseRespondInterface
}
