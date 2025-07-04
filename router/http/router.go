package httpRouter

import (
	"github.com/gin-gonic/gin"

	"silver-train/router/http/controller/auth"
)

func Run() {
	r:= gin.Default()

	jwtRouterGroup := r.Group("jwt")
	{
		jwtRouterGroup.GET("/get", authController.Get)
		jwtRouterGroup.PUT("/refresh", authController.Refresh)
	}

	userRouterGroup := r.Group("user")
	{
		userRouterGroup.GET("/me", authController.Current)
		userRouterGroup.GET("/logout", authController.Revoke)
	}

	r.Run()
}
