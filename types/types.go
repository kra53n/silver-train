package types

type AccessToken string

type RefreshToken string

type RefreshTokenDB string

type ErrorResponse struct {
	Error string `json:"error"`
}
