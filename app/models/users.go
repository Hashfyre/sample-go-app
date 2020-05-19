package models

import (
	"log"

	"github.com/jinzhu/gorm"

	db "github.com/hashfyre/sample-go-app/app/database"
)

// User defines the attributes of a real world user entity who either signs up
// on the platform or is invited to the same
type User struct {
	BaseModel
	FirstName string `gorm:"varchar(255);not null"`
	LastName  string `gorm:"varchar(255);not null"`
	Email     string `gorm:"column:email;unique_index"`
	Password  string `gorm:"varchar(255);column:password;not null"`
}

// CreateOne inserts a new user in the DB
func (user User) CreateOne() (*User, error) {
	database := db.GetDB()

	count, err := user.Count()
	if err != nil {
		log.Println(err)
		return nil, errCreateUser
	}

	if count != 0 {
		return nil, errAlreadyExists
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		return nil, errCreateUser
	}

	return &user, nil
}

// GetOne returns one User entity for a given condition
func (user User) GetOne() (*User, error) {
	database := db.GetDB()
	var result User

	count, err := user.Count()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if count != 1 {
		log.Println(errGetUser)
		return nil, errGetUser
	}

	err = database.DB.Where(user).First(&result).Error
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return nil, errRecordNotFound
		}

		return nil, errGetUser
	}

	return &result, nil
}

// GetAny gets Users for a given condition
func (user User) GetAny() ([]User, error) {
	database := db.GetDB()
	var result []User

	err := database.DB.Where(user).Find(&result).Error
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errGetUsers
	}

	return result, nil
}

// Count gets the number of users for a condition
func (user User) Count() (int, error) {
	database := db.GetDB()
	var count int
	err := database.DB.Model(&User{}).Where(user).Count(&count).Error
	if err != nil {
		return -1, errCountUser
	}

	return count, nil
}

// UpdateOne updates a given User entity that satisfies the condition
// You could update properties of an User to database returning with error info.
//  err := db.Model(userModel).Update(User{Username: "wangzitian0"}).Error
func (user User) UpdateOne(data User) (*User, error) {
	database := db.GetDB()

	result, err := user.GetOne()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = database.DB.Model(result).Where(user).Updates(data).Error
	if err != nil {
		log.Println(err)
		return nil, errUpdateUser
	}

	return result, nil
}

// ---------------------------------------------------------------------
