package router

import (
	"web-lab2/platform/database"
	"web-lab2/platform/middleware"
	"web-lab2/web/app/csrfvan"
	"web-lab2/web/app/home"
	"web-lab2/web/app/sqlvan"
	"web-lab2/web/app/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(db *database.PostgresStorage, csrfMid *middleware.CSRFMiddleware) *gin.Engine {
	rtr := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	rtr.Use(sessions.Sessions("mysession", store))

	rtr.Static("/public", "web/static")
	rtr.LoadHTMLGlob("web/templates/*")

	rtr.GET("/", home.Page(db))

	rtr.GET("/user/signup", user.SignupPage)
	rtr.GET("/user/login", user.LoginPage)
	rtr.POST("/user/signup", user.Signup(db))
	rtr.POST("/user/login", user.Login(db))
	rtr.DELETE("/user/delete/off", middleware.Auth(), user.Delete(db))
	rtr.DELETE("/user/delete/on", middleware.Auth(), csrfMid.ValidateCSRF(), user.Delete(db))
	rtr.DELETE("/user/logout", middleware.Auth(), user.Logout)
	rtr.GET("/csrf", csrfvan.Page(csrfMid, db))

	rtr.GET("/sql", sqlvan.Page)
	rtr.POST("/sql", sqlvan.SimLogin(db))

	return rtr
}
