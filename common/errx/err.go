package errx

import "errors"

var (
	ErrTokenIsRequired = errors.New("token is required")
	ErrClientNotFound  = errors.New("client is not found")
)

func CheckToken(token string) bool {
	if token == "" {
		return true
	}
	return false
}
