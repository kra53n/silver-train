package authController

import (
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

// @Summary Get tokens
// @Description Get access and refresh tokens
// @Tags auth
// @Accept  json
// @Produce  json
// @Param guid query string true "User GUID"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/jwt/get [get]
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

// @Summary Refresh tokens
// @Description Exchange a valid refresh token for new access and refresh tokens
// @Tags auth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Access-Token header string true "Current access token"
// @Param Refresh-Token header string true "Valid refresh token to exchange"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/jwt/refresh [put]
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
