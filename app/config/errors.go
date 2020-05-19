package config

import (
	"errors"
)

var packageName = "config"

//ErrNotFoundConfigVar -  Occurs when a needed environment variable is not set the the OS env
var (
	ErrNotFoundConfigVar = errors.New("Missing environment variable")
)
