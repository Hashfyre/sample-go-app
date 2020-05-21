package config

import (
	"log"
)

// Config represents a set of environment variables nimbus needs to bootstrap
type Config struct {
	Tracer   JaegarTracer
	Database Database
	Port     string
}

// global config singleton
var config *Config

// Init Looks up all needed environment variables and stores them in Config
func Init() (*Config, error) {
	appPort, err := configCheck("APP_PORT")
	if err != nil {
		return nil, err
	}

	database, err := initDatabase()
	if err != nil {
		return nil, err
	}

	tracer, err := initJaegerTracer()
	if err != nil {
		return nil, err
	}

	log.Println("Initialized config")
	config = &Config{
		Database: *database,
		Tracer:   *tracer,
		Port:     appPort,
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
