package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/util"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		// token := ctx.Request.Header.Get("Authorization")
		token := ctx.Query("token")
		if token == "" {
			code = http.StatusUnauthorized
			data = gin.H{
				"code": code,
				"msg":  http.StatusText(code),
			}
			ctx.JSON(code, data)
			ctx.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			data = gin.H{
				"code": code,
				"msg":  http.StatusText(code),
			}
			ctx.JSON(code, data)
			ctx.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusUnauthorized
			data = gin.H{
				"code": code,
				"msg":  http.StatusText(code),
			}
			ctx.JSON(code, data)
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
