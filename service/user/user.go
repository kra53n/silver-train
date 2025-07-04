package userService

import (
	"silver-train/types"
	"silver-train/util"
	"silver-train/service/auth"
)

func Me(access types.AccessToken) (string, error) {
	claims, err := util.ParseAccessToken(access)
	if err != nil {
		return "", err
	}
	return claims["sub"].(string), nil
}

func Logout(access types.AccessToken) error {
	claims, err := util.ParseAccessToken(access)
	if err != nil {
		return err
	}
	return authService.RevokeAll(claims["sub"].(string))
}
