package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Postgres represents database connection parameters
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

// GORM represents config needed for database connection parameter customizations
type GORM struct {
	Log             bool
	MaxConnLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

type Database struct {
	GORM     GORM
	Postgres Postgres
}

// Config represents a set of environment variables nimbus needs to bootstrap
type Config struct {
	Database Database
	Port     string
}

// global config singleton
var config *Config

// Init Looks up all needed environment variables and stores them in Config
func Init() (*Config, error) {
	pgHost, err := configCheck("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}

	pgPort, err := configCheck("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}

	dbPgPort, err := strconv.Atoi(pgPort)
	if err != nil {
		return nil, err
	}

	pgUser, err := configCheck("POSTGRES_USER")
	if err != nil {
		return nil, err
	}

	pgPassword, err := configCheck("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}

	pgDB, err := configCheck("POSTGRES_DB")
	if err != nil {
		return nil, err
	}

	gormLog, err := configCheck("GORM_LOG")
	if err != nil {
		return nil, err
	}

	dbGormLog, err := strconv.ParseBool(gormLog)
	if err != nil {
		return nil, err
	}

	gormMaxConnLifetime, err := configCheck("GORM_CONN_MAX_LIFETIME")
	if err != nil {
		return nil, err
	}

	dbGormMaxConnLifetime, err := time.ParseDuration(gormMaxConnLifetime)
	if err != nil {
		return nil, err
	}

	gormMaxOpenConn, err := configCheck("GORM_CONN_MAX_OPEN")
	if err != nil {
		return nil, err
	}

	dbGormMaxOpenConn, err := strconv.Atoi(gormMaxOpenConn)
	if err != nil {
		return nil, err
	}

	gormMaxIdleConn, err := configCheck("GORM_CONN_MAX_IDLE")
	if err != nil {
		return nil, err
	}

	dbGormMaxIdleConn, err := strconv.Atoi(gormMaxIdleConn)
	if err != nil {
		return nil, err
	}

	appPort, err := configCheck("APP_PORT")
	if err != nil {
		return nil, err
	}

	log.Println("Initialized config")
	config = &Config{
		Database: Database{
			GORM: GORM{
				Log:             dbGormLog,
				MaxConnLifetime: dbGormMaxConnLifetime,
				MaxOpenConns:    dbGormMaxOpenConn,
				MaxIdleConns:    dbGormMaxIdleConn,
			},
			Postgres: Postgres{
				Host:     pgHost,
				Port:     dbPgPort,
				User:     pgUser,
				Password: pgPassword,
				DB:       pgDB,
			},
		},
		Port: appPort,
	}
	return config, nil
}

// Get helps get the singleton config object
func Get() (*Config, error) {
	if config == nil {
		return Init()
	}
	return config, nil
}

//Envcheck checks if the env variable is declared
func configCheck(configVar string) (string, error) {
	value, ok := os.LookupEnv(configVar)
	if !ok {
		log.Printf("%v - %s", ErrNotFoundConfigVar, configVar)
		return "", ErrNotFoundConfigVar
	}

	return value, nil
}
