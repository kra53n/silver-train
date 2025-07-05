package util

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	jwt "github.com/dgrijalva/jwt-go"

	"silver-train/vars"
	"silver-train/types"
)

func GenerateAccessToken(guid string) (types.AccessToken, uuid.UUID, error) {
	id := uuid.New()
	token := jwt.NewWithClaims(vars.SigningMethod, jwt.StandardClaims{
		Id: id.String(),
		Subject: guid,
		ExpiresAt: time.Now().Add(vars.AccessTokenExpire).Unix(),
	})
	s, err := token.SignedString([]byte(vars.JwtSecret()))
	return types.AccessToken(s), id, err
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

func ParseAccessToken(access types.AccessToken) (jwt.MapClaims, error) {
	token, err := jwt.Parse(string(access), func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != vars.SigningMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(vars.JwtSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
