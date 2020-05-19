package models

import (
	db "github.com/hashfyre/sample-go-app/app/database"
)

// Migrate - applies schema migrations for each model
func Migrate() {
	database := db.GetDB()
	database.DB.AutoMigrate(&User{})
}
