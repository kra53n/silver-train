package userService

import (
	"silver-train/types"
	"silver-train/util"
)

func Me(access types.AccessToken) (string, error) {
	claims, err := util.ParseAccessToken(access)
	if err != nil {
		return "", err
	}
	return claims["sub"].(string), nil
}
