package header

import (
	"errors"
	"log"
	"strings"
)

// ErrInvalidTokenType occurs when authorization header does not have the format
// basic <token>
var (
	ErrEmptyHeaderValue = errors.New("Authorization header has no value")
	ErrInvalidToken     = errors.New("Token header is invalid")
	ErrInvalidTokenType = errors.New("Token header has invalid token type")
)

// ParseAuth - checks whether the given token is prefixed correctly
// TODO: Find a place this generic function
func ParseAuth(prefix string, token string) (string, error) {
	if len(token) == 0 {
		log.Print(ErrEmptyHeaderValue)
		return "", ErrEmptyHeaderValue
	}

	token = strings.TrimSuffix(token, " ")
	tokenParts := strings.Split(token, " ")

	if len(tokenParts) != 2 {
		log.Print(ErrInvalidToken)
		return "", ErrInvalidToken
	}

	tokenType := tokenParts[0]
	if strings.ToLower(tokenType) != prefix {
		log.Print(ErrInvalidTokenType)
		return "", ErrInvalidTokenType
	}

	return tokenParts[1], nil
}
