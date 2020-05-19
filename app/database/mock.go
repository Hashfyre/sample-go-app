package database

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Needed for gorm postgres ops

	dbpg "github.com/hashfyre/sample-go-app/pkg/database/postgres"
)

// SetMockDB - sets mock connection params to the global DB singleton
// for testing purposes
func SetMockDB() sqlmock.Sqlmock {
	db, mock, err := dbpg.Mock()
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	return *mock
}
