package constant

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenExpire = 15 * time.Minute
	RefreshTokenExpire = 7 * 24 * time.Hour
)

var (
	SigningMethod = jwt.SigningMethodHS256
)
