package config

import (
	"log"
	"os"
)

//Envcheck checks if the env variable is declared
func configCheck(configVar string) (string, error) {
	value, ok := os.LookupEnv(configVar)
	if !ok {
		log.Printf("%v - %s", ErrNotFoundConfigVar, configVar)
		return "", ErrNotFoundConfigVar
	}

	return value, nil
}
