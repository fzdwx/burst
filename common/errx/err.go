package errx

import "errors"

var (
	TokenIsRequired = errors.New("token is required")
)

func CheckToken(token string) bool {
	if token == "" {
		return true
	}
	return false
}
