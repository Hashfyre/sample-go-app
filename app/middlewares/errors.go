package middleware

import (
	"errors"
)

var (
	errHostNameUnset       = errors.New("Hostname not set")
	errExpiredAccess       = errors.New("Access has expired")
	errUserUnauthorized    = errors.New("You are not authorized to perform this action")
	errMalformedAuthHeader = errors.New("Malformed auth header")
)
