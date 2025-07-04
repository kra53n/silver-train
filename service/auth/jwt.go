package authService

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	
	"silver-train/constant"
	"silver-train/util"
	"silver-train/types"
	"silver-train/db"
	"silver-train/model/auth"
)


func GetTokens(guid, userAgent, ipAddress string) (types.AccessToken, types.RefreshToken, error) {
	access, err := util.GenerateAccessToken(guid)
	if err != nil {
		return "", "", err
	}
	refresh, refreshDB, err := util.GenerateRefreshToken(guid)
	if err != nil {
		return "", "", err
	}
	res := db.DB.Create(&authModel.RefreshToken{
		UserID:  guid,
		TokenHash: string(refreshDB),
		UserAgent: userAgent,
		IPAddress: ipAddress,
		ExpiresAt: time.Now().Add(constant.RefreshTokenExpire),
		Revoked:   false,
	})
	if res.Error != nil {
		return "", "", res.Error
	}
	return access, refresh, nil
}

func Refresh(access types.AccessToken, refresh types.RefreshToken, userAgent string, ipAddress string) (types.AccessToken, types.RefreshToken, error) {
	claims, err := util.ParseAccessToken(access)
	if err != nil {
		return types.AccessToken(""), types.RefreshToken(""), err
	}
	userId := claims["sub"].(string)
	tokenModels := []authModel.RefreshToken{}
	db.DB.Where("user_id = ? and revoked = 0", userId).Find(&tokenModels)

	var createNewTokens bool

	for _, tokenModel := range tokenModels {
		db.DB.Model(&tokenModel).Update("revoked", true)
		err = bcrypt.CompareHashAndPassword([]byte(tokenModel.TokenHash), []byte(refresh))
		if err != nil {
			continue
		}
		if userAgent == tokenModel.UserAgent {
			if ipAddress == tokenModel.IPAddress {
				defer func () {
					go util.SendMsgAtWebHook("the user ip address has been change")
				}()
			}
			createNewTokens = true
		}
	}

	if createNewTokens {
		return GetTokens(userId, userAgent, ipAddress)
	}
	return types.AccessToken(""), types.RefreshToken(""), fmt.Errorf("something went wrong")
}

func RevokeAll(userId string) error {
	tokenModels := []authModel.RefreshToken{}
	db.DB.Where("user_id = ? and revoked = 0", userId).Find(&tokenModels)
	for _, tokenModel := range tokenModels {
		db.DB.Model(&tokenModel).Update("revoked", true)
	}
	return nil
}
