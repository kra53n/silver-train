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
		jwtRouterGroup.GET("/current", authController.Current)
		jwtRouterGroup.GET("/revoke", authController.Revoke)
	}
	r.Run()
}
