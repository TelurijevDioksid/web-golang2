package home

import (
	"web-lab2/platform/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Page(db *database.PostgresStorage) gin.HandlerFunc {
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

		ctx.HTML(200, "home.tmpl", gin.H{
			"User": username,
		})
	}
}
