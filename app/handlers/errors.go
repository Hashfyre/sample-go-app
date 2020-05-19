package handlers

import (
	"github.com/hashfyre/sample-go-app/app/types"
)

var errUserInfoContext = types.ResponseError{
	Code:    "ERR_BAD_REQUEST",
	Message: "Invalid user info context",
}
