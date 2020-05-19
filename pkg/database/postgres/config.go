package postgres

import (
	"time"
)

type Config struct {
	Postgres Postgres
	GORM     GORM
}

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
