package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if sessions.Default(ctx).Get("userID") == nil {
			ctx.AbortWithStatus(401)
		}
		ctx.Next()
	}
}
