package authService

import (
	"silver-train/util"
	"silver-train/types"
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
	return access, refresh, nil
}
