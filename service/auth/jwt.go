package authService

import (
	"fmt"
	"time"

	// "golang.org/x/crypto/bcrypt"
	
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
		fmt.Println("Some err occured:", err)
		return types.AccessToken("s"), types.RefreshToken("2"), err
	}
	userId := claims["sub"].(string)
	tokenModels := []authModel.RefreshToken{}
	db.DB.Where("user_id = ? and revoked = 0", userId).First(&tokenModels)

	createNewTokens := true

	for _, tokenModel := range tokenModels {
		// TODO: fix
		// err = bcrypt.CompareHashAndPassword([]byte(tokenModel.TokenHash), []byte(refresh))
		// if err != nil {
		// 	return types.AccessToken(""), types.RefreshToken(""), err
		// }
		if userAgent != tokenModel.UserAgent || ipAddress != tokenModel.IPAddress {
			fmt.Println(userAgent, tokenModel.UserAgent)
			fmt.Println(ipAddress, tokenModel.IPAddress)
			createNewTokens = false
		}

		db.DB.Model(&tokenModel).Update("revoked", true)
	}

	if createNewTokens {
		return GetTokens(userId, userAgent, ipAddress)
	}
	return types.AccessToken(""), types.RefreshToken(""), fmt.Errorf("bad")
}
