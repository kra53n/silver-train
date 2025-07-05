package vars

import (
	"os"
)

func JwtSecret() string {
	res := os.Getenv("JWT_SECRET")
	if res == "" {
		return jwtSecret
	}
	return res
}

func MsgWebhook() string {
	res := os.Getenv("MSG_WEBHOOK")
	if res == "" {
		return msgWebhook
	}
	return res
}
