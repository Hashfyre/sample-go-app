package main

import (
	"log"

	"github.com/hashfyre/sample-go-app/app/config"
	db "github.com/hashfyre/sample-go-app/app/database"
	"github.com/hashfyre/sample-go-app/app/models"
	"github.com/hashfyre/sample-go-app/app/routers"
	//"github.com/hashfyre/sample-go-app/pkg/trace"
)

// @title  sample-go-app
// @version 0.1
// @description A sample microservice built using gin, gorm
// @host localhost:8080
// @BasePath /

// Ignore all the logging for now mismatch in compatibility
func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Println(err)
		panic("unable to initialize config")
	}

	database := db.GetDB()
	defer func() {
		err := database.DB.Close()
		if err != nil {
			panic(err)
		}
	}()

	models.Migrate()

	app := routers.Setup()
	err = app.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
