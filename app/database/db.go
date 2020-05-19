package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // Needed for gorm postgres ops
	"log"

	"github.com/hashfyre/sample-go-app/app/config"
	pgdb "github.com/hashfyre/sample-go-app/pkg/database/postgres"
)

// DB - global DB object to be used as a singleton
var DB *pgdb.Database

// InitDB - sets the connection parameters for the global DB singleton
func InitDB() (*pgdb.Database, error) {
	cfg, err := config.Get()
	if err != nil {
		return nil, err
	}

	pgConf := pgdb.Config{
		Postgres: pgdb.Postgres{
			Host:     cfg.Database.Postgres.Host,
			Port:     cfg.Database.Postgres.Port,
			User:     cfg.Database.Postgres.User,
			Password: cfg.Database.Postgres.Password,
			DB:       cfg.Database.Postgres.DB,
		},
		GORM: pgdb.GORM{
			Log:             cfg.Database.GORM.Log,
			MaxConnLifetime: cfg.Database.GORM.MaxConnLifetime,
			MaxOpenConns:    cfg.Database.GORM.MaxOpenConns,
			MaxIdleConns:    cfg.Database.GORM.MaxIdleConns,
		},
	}

	return pgdb.Init(pgConf), nil
}

// GetDB - implments the singleton patter by returning the global DB object
func GetDB() *pgdb.Database {
	if DB == nil {
		var err error
		DB, err = InitDB()
		if err != nil {
			log.Fatal("Unable to initialize database connection")
			return nil
		}

	}

	return DB
}
