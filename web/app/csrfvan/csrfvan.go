package csrfvan

import (
	"web-lab2/platform/middleware"

	"github.com/gin-gonic/gin"
)

func Page(cs *middleware.CSRFMiddleware) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.HTML(200, "csrfvan.tmpl", gin.H{
            "Token": cs.GetToken(),
        })
    }
}
