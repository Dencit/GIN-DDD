package midware

import (
	BaseErr "app/base/err"
	"app/extend/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
notes: 中间件基础
*/

const (
	TokenSalt = "token"
)

func UsrAuth() gin.HandlerFunc {
	return func(context *gin.Context) {

		token := context.Request.Header.Get("Authorization") // 访问令牌
		//log.Println(token)

		//todo:: 临时生成token
		//log.Println("Bearer " + BaseHelper.MD5([]byte(TokenSalt)))

		if strings.ToLower("Bearer "+encrypt.MD5([]byte(TokenSalt))) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			context.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			context.Abort()
			//
			Code, Message := BaseErr.Root("AUTH_MUST")
			ErrData := gin.H{"code": Code, "message": Message}
			context.AbortWithStatusJSON(http.StatusUnauthorized, ErrData)
			return
		}
	}

}
