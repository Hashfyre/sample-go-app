package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// EncryptString encrypts a given plaintext password string before storage into database
func EncryptString(password string) ([]byte, error) {
	if len(password) == 0 {
		return nil, errors.New("password should not be empty")
	}

	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return passwordHash, nil
}
