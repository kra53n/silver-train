package vars

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenExpire = 15 * time.Minute
	RefreshTokenExpire = 7 * 24 * time.Hour
)

var (
	SigningMethod = jwt.SigningMethodHS512
)

// application will use values for environment variables that was not found
const (
	jwtSecret = "a-string-secret-at-least-256-bits-long"
	msgWebhook = "http://127.0.0.1:2020"
)
