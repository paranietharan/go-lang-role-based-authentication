package router

import (
	"go-lang-role-based-authentication/pkg/controller"
	"go-lang-role-based-authentication/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.POST("/signup", controller.SignUp())
	router.POST("/login", controller.Login())

	protected := router.Group("/")
	protected.Use(middleware.Authenticate())
	protected.GET("/users", controller.GetUsers())
	protected.GET("/users/:user_id", controller.GetUser())

	router.Run(":8080")
}
