package util

import (
	"time"
	"math/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"

	"silver-train/constant"
	"silver-train/types"
)

func GenerateAccessToken(guid string) (types.AccessToken, error) {
	// TODO: use SHA512 algo for sign
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: guid,
		ExpiresAt: time.Now().Add(constant.AccessTokenExpire).Unix(),
	})
	s, err := token.SignedString([]byte("TODO: make sugar string"))
	return types.AccessToken(s), err
}

func GenerateRefreshToken(guid string) (types.RefreshToken, types.RefreshTokenDB, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", "", err
	}
	tokenString := base64.URLEncoding.EncodeToString(token)
	hash, err := bcrypt.GenerateFromPassword([]byte(tokenString), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return types.RefreshToken(tokenString), types.RefreshTokenDB(hash), nil
}
