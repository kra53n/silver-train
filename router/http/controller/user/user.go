package userController

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"silver-train/service/user"
	"silver-train/types"
)

type MeResponse struct {
    GUID string `json:"guid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type LogoutResponse struct {
    Success bool `json:"success" example:"true"`
}

// @Summary Get current user info
// @Description Returns GUID of the authenticated user
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Access-Token header string true "Valid access token"
// @Success 200 {object} MeResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/user/me [get]
func Me(c *gin.Context) {
	access := types.AccessToken(c.GetHeader("Access-Token"))
	userId, err := userService.Me(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, MeResponse{GUID: userId})
}

// @Summary Logout user
// @Description Invalidates the current session and tokens
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Access-Token header string true "Valid access token to invalidate"
// @Success 200 {object} LogoutResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/user/logout [get]
func Logout(c *gin.Context) {
	access := types.AccessToken(c.GetHeader("Access-Token"))
	err := userService.Logout(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, LogoutResponse{Success: true})
}
