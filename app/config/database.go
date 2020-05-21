package config

import (
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

// Database represents config for DB connector + ORM
type Database struct {
	GORM     GORM
	Postgres Postgres
}

func initPostgres() (*Postgres, error) {
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

	return &Postgres{
		Host:     pgHost,
		Port:     dbPgPort,
		User:     pgUser,
		Password: pgPassword,
		DB:       pgDB,
	}, nil
}

func initGORM() (*GORM, error) {
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

	return &GORM{
		Log:             dbGormLog,
		MaxConnLifetime: dbGormMaxConnLifetime,
		MaxOpenConns:    dbGormMaxOpenConn,
		MaxIdleConns:    dbGormMaxIdleConn,
	}, nil
}

func initDatabase() (*Database, error) {
	pg, err := initPostgres()
	if err != nil {
		return nil, err
	}

	gorm, err := initGORM()
	if err != nil {
		return nil, err
	}

	return &Database{
		Postgres: *pg,
		GORM:     *gorm,
	}, nil
}
