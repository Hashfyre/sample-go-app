package postgres

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Needed for gorm postgres ops

	p_uuid "github.com/hashfyre/sample-go-app/pkg/uuid"
)

// Mock - sets mock connection params to the global DB singleton
// for testing purposes
func Mock() (*Database, *sqlmock.Sqlmock, error) {
	randID := p_uuid.MustUUID()
	_, mock, err := sqlmock.NewWithDSN("sqlmock_db_" + randID.String())
	if err != nil {
		log.Println("Error setting mock_sql database : ", err)
		return nil, nil, err
	}

	db, err := gorm.Open("sqlmock", "sqlmock_db_"+randID.String())
	if err != nil {
		log.Println("Error connecting to database : ", err)
		return nil, nil, err
	}

	return &Database{DB: db}, &mock, nil
}
