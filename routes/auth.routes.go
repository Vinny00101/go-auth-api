package routes

import (
	auth_controllers "go-api/controllers"
	"go-api/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup_routes_auth(server *gin.RouterGroup) {
	authController := &auth_controllers.AuthController{}

	auth := server.Group("/auth")
	{
		auth.POST("/register", authController.Register_auth)
		auth.POST("/login", authController.Login_auth)
		auth.GET("/me", middlewares.AuthMiddleware(), authController.Me_auth)
	}
}
