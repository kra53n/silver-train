package authController

import (
	_ "fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"

	"silver-train/service/auth"
	"silver-train/types"
)

type AuthRequest struct {
	GUID string `form:"guid" binding:"required,uuid"`
}

type TokenResponse struct {
	Access types.AccessToken `json:"access"`
	Refresh types.RefreshToken `json:"refresh"`
}

func Get(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guid := c.Query("guid")
	userAgent := c.GetHeader("User-Agent")
	ipAddress := c.ClientIP()
	access, refresh, err := authService.GetTokens(guid, userAgent, ipAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, TokenResponse{
		Access: access,
		Refresh: refresh,
	})
}

func Refresh(c *gin.Context) {
	var access types.AccessToken
	var refresh types.RefreshToken
	var err error
	access = types.AccessToken(c.GetHeader("Access-Token"))
	refresh = types.RefreshToken(c.GetHeader("Refresh-Token"))
	userAgent := c.GetHeader("User-Agent")
	ipAddress := c.ClientIP()
	access, refresh, err = authService.Refresh(access, refresh, userAgent, ipAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, TokenResponse{
		Access: access,
		Refresh: refresh,
	})
}

func Current(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func Revoke(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
