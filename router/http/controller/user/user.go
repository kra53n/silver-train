package userController

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"silver-train/service/user"
	"silver-train/types"
)

func Me(c *gin.Context) {
	access := types.AccessToken(c.GetHeader("Access-Token"))
	userId, err := userService.Me(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"guid": userId})
}

func Logout(c *gin.Context) {
	access := types.AccessToken(c.GetHeader("Access-Token"))
	err := userService.Logout(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
