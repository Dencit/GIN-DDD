package exception

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
notes: 异常输出类
*/

//声明结构

type ErrStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//应用异常

func App(context *gin.Context, Code int, Message string) {
	ErrData := ErrStruct{
		Code:    Code,
		Message: Message,
	}
	context.AbortWithStatusJSON(http.StatusInternalServerError, ErrData)
	log.Println(ErrData)
	panic(ErrData)
}
