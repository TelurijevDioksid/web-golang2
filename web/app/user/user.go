package user

import (
	"web-lab2/platform/database"
	"web-lab2/platform/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Signup(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := new(models.UserDto)
		if err := ctx.Bind(user); err != nil {
			ctx.JSON(400, "Invalid user data")
			return
		}

		if user.Username == "" || user.Password == "" {
			ctx.JSON(400, "Username and password are required")
			return
		}

		if db.UsernameExists(user.Username) == true {
			ctx.JSON(400, "Username already exists")
			return
		}

		id, err := db.CreateUser(user)
		if err != nil {
			ctx.JSON(500, "Error creating user")
			return
		}

		session := sessions.Default(ctx)
		session.Set("userID", id)
		session.Save()

		ctx.JSON(200, "User created")
	}
}

func SignupPage(ctx *gin.Context) {
	ctx.HTML(200, "signup.tmpl", nil)
}

func Login(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userDto := new(models.UserDto)
		if err := ctx.Bind(userDto); err != nil {
			ctx.JSON(400, "Invalid user data")
			return
		}

		if userDto.Username == "" || userDto.Password == "" {
			ctx.JSON(400, "Username and password are required")
			return
		}

		user, err := db.GetUser(userDto)
		if err != nil {
			ctx.JSON(400, "Invalid password or username")
			return
		}

		session := sessions.Default(ctx)
		session.Set("userID", user.ID)
		session.Save()

		ctx.JSON(200, "Logged in")
	}
}

func LoginPage(ctx *gin.Context) {
	ctx.HTML(200, "login.tmpl", nil)
}

func Delete(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		id := session.Get("userID")
		if id == nil {
			ctx.JSON(400, "You are not logged in")
			return
		}

		err := db.DeleteUser(id.(string))
		if err != nil {
			ctx.JSON(500, "Error deleting user")
			return
		}

		session.Clear()
		session.Save()

        ctx.JSON(200, "User deleted")
	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(302, "/")
}
