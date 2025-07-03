package authService

import (
	"time"
	
	"silver-train/constant"
	"silver-train/util"
	"silver-train/types"
	"silver-train/db"
	"silver-train/model/auth"
)


func GetTokens(guid string) (types.AccessToken, types.RefreshToken, error) {
	access, err := util.GenerateAccessToken(guid)
	if err != nil {
		return "", "", err
	}
	refresh, refreshDB, err := util.GenerateRefreshToken(guid)
	if err != nil {
		return "", "", err
	}
	_ = refreshDB
	res := db.DB.Create(&authModel.RefreshToken{
		UserGUID:  guid,
		TokenHash: string(refreshDB),
		// UserAgent: userAgent, // TODO
		// IPAddress: ipAddress, // TODO
		ExpiresAt: time.Now().Add(constant.RefreshTokenExpire),
		Revoked:   false,
	})
	if res.Error != nil {
		return "", "", res.Error
	}
	return access, refresh, nil
}
