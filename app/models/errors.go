package models

import (
	"errors"
)

var (
	// Generic errors
	errAlreadyExists  = errors.New("record already exists")
	errRecordNotFound = errors.New("record does not exist")
	errNotOne         = errors.New("condition matches either zero or more than one records")

	// User specific errors
	errCreateUser = errors.New("failed to create user")
	errGetUser    = errors.New("failed to get user")
	errGetUsers   = errors.New("failed to get list of users")
	errCountUser  = errors.New("failed to get count of user")
	errUpdateUser = errors.New("failed to update user")
)
