package controllers

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/hashfyre/sample-go-app/app/models"
	"github.com/hashfyre/sample-go-app/pkg/crypto"
	p_uuid "github.com/hashfyre/sample-go-app/pkg/uuid"
)

// UsersRegistration registers a new user
// TODO: return model.user from this
func UsersRegistration(firstName, lastName, email, password string) (*models.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Unable to hash password using bcrypt")
	}

	baseModel := models.BaseModel{ID: p_uuid.MustUUID()}

	user, err := models.User{
		BaseModel: baseModel,
		Password:  string(passwordHash),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}.CreateOne()
	if err != nil {
		log.Print(err)
		return nil, errCreateUser
	}

	return user, nil
}

// UserLogin login user
func UserLogin(email string, password string) (*models.User, error) {
	user, err := models.User{Email: email}.GetOne()
	if err != nil {
		return nil, errors.New("User does not exist")
	}

	if isValidPassword(*user, password) != nil {
		return nil, errors.New("Invalid credentials")
	}

	return user, nil
}

// GetAllUsers returns all user data across the platform
// sans sensitive data like password
func GetAllUsers() ([]models.User, error) {
	users, err := models.User{}.GetAny()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// GetUser get basic info of user
func GetUser(userID uuid.UUID) (*models.User, error) {
	user, err := models.User{
		BaseModel: models.BaseModel{
			ID: userID,
		}}.GetOne()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// UpdateUser update user's basic info
func UpdateUser(userID uuid.UUID, fields models.User) (*models.User, error) {
	user, err := models.User{
		BaseModel: models.BaseModel{
			ID: userID,
		},
	}.UpdateOne(fields)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// UpdatePassword updates user password
// What is the matter with these int64 returns, goddamnit!
func UpdatePassword(userID uuid.UUID, currentPassword string, newPassword string) (int64, error) {
	user, err := models.User{BaseModel: models.BaseModel{ID: userID}}.GetOne()
	if err != nil {
		return 0, err
	}

	err = isValidPassword(*user, currentPassword)
	if err != nil {
		return 0, err
	}

	newPasswordHashByte, err := crypto.EncryptString(newPassword)
	newPasswordHash := string(newPasswordHashByte)
	if err != nil {
		return 0, err
	}

	_, dbErr := models.User{BaseModel: models.BaseModel{ID: userID}}.UpdateOne(models.User{
		Password: newPasswordHash,
	})
	if dbErr != nil {
		log.Println(dbErr)
		return 0, dbErr
	}

	return 1, nil
}

// DeleteUser disables current user
// Another int64!
func DeleteUser(userID uuid.UUID) (int64, error) {
	now := time.Now()

	_, err := models.User{BaseModel: models.BaseModel{ID: userID}}.UpdateOne(
		models.User{
			BaseModel: models.BaseModel{
				DeletedAt: &now,
			},
		})
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return 1, err
}

// IsValidPassword checks whether the stored password is valid for an User
// Database will only save the encrypted string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
// Should be in controller
func isValidPassword(user models.User, password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
