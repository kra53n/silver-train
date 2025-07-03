package constant

import (
	"time"
)

const (
	AccessTokenExpire = 15 * time.Minute
	RefreshTokenExpire = 7 * 24 * time.Hour
)
