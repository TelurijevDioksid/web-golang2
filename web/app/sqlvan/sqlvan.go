package sqlvan

import (
	"web-lab2/platform/database"
	"web-lab2/platform/models"

	"github.com/gin-gonic/gin"
)

func Page(ctx *gin.Context) {
	ctx.HTML(200, "sqlvan.tmpl", nil)
}

func SimLogin(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userDto := new(models.SqlInjUser)
		if err := ctx.Bind(userDto); err != nil {
			ctx.JSON(400, "Invalid user data")
			return
		}

		if userDto.Username == "" || userDto.Password == "" {
			ctx.JSON(400, "Username and password are required")
			return
		}

		if userDto.Inj == "true" {
			rows, err := db.SimGetUser(userDto)
			if err != nil {
				ctx.JSON(400, "Invalid password or username")
				return
			}

			res := make([]models.UserDto, 0)
			for rows.Next() {
				id := ""
				user := new(models.UserDto)
				err := rows.Scan(&id, &user.Username, &user.Password)
				if err != nil {
					ctx.JSON(500, "Error scanning user")
					return
				}
				res = append(res, *user)
			}

			ctx.JSON(200, res)
		} else {
			user, err := db.GetUser(&models.UserDto{
				Username: userDto.Username,
				Password: userDto.Password,
			})
			if err != nil {
				ctx.JSON(400, "Invalid password or username")
				return
			}

			ctx.JSON(200, user)
		}
	}
}
