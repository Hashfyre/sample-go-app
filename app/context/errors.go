package context

import "errors"

var (
	errMissingContextKey = errors.New("Context is missing key")
	errBadContextKey     = errors.New("Malformed context key value")
)
