package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect for gorm
)

// Database - wraps a gorm.DB pointer
// TODO: verify if this is being used, else delete
type Database struct {
	DB *gorm.DB
}

// Init - sets the connection parameters for the global DB singleton
func Init(config Config) *Database {
	log.Print("database conf: ", config)
	dsn := getDSN(config.Postgres)
	log.Println("dsn: ", dsn)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Println("error connecting to database : ", err)
		os.Exit(-1)
	}

	// Setting this here and not from config
	db.DB().SetMaxIdleConns(config.GORM.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.GORM.MaxOpenConns)
	db.DB().SetConnMaxLifetime(config.GORM.MaxConnLifetime)
	db.LogMode(config.GORM.Log)

	return &Database{
		DB: db,
	}
}

// DSN generates the postgres connection string
func getDSN(postgres Postgres) string {
	// Remove SSL mode off before production
	log.Println("postgres conf: ", postgres)
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", postgres.Host, postgres.Port, postgres.User, postgres.DB, postgres.Password)
}
