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

	api := r.Group("api")

	v1 := api.Group("v1")
	{
		jwt := v1.Group("jwt")
		{
			jwt.GET("/get", authController.Get)
			jwt.PUT("/refresh", authController.Refresh)
		}

		user := v1.Group("user")
		{
			user.GET("/me", userController.Me)
			user.GET("/logout", userController.Logout)
		}
	}


	r.Run()
}
