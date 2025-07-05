package httpRouter

import (
	"github.com/gin-gonic/gin"

	_ "silver-train/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"silver-train/router/http/controller/auth"
	"silver-train/router/http/controller/user"
)

// @title Auth Service API
// @version 1.0
// @description JWT Authentication Service
// @host localhost:8080
// @BasePath /api/v1
func Run() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtRouterGroup := r.Group("jwt")
	{
		jwtRouterGroup.GET("/get", authController.Get)
		jwtRouterGroup.PUT("/refresh", authController.Refresh)
	}

	userRouterGroup := r.Group("user")
	{
		userRouterGroup.GET("/me", userController.Me)
		userRouterGroup.GET("/logout", userController.Logout)
	}

	r.Run()
}
