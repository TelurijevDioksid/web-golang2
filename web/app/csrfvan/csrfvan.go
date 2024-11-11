package csrfvan

import (
	"web-lab2/platform/database"
	"web-lab2/platform/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Page(cs *middleware.CSRFMiddleware, db *database.PostgresStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
		username := ""
		session := sessions.Default(ctx)
		id := session.Get("userID")
		if id != nil {
			user, err := db.GetUserByID(id.(string))
			if err == nil {
				username = user.Username
			}
        }

        ctx.HTML(200, "csrfvan.tmpl", gin.H{
            "Token": cs.GetToken(),
            "User": username,
        })
    }
}
