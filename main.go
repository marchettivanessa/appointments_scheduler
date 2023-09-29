package main

import (
	"appointments_scheduler/api"
	"appointments_scheduler/config"
	"appointments_scheduler/database"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	//Start config
	c := config.MustParseConfig()

	// Iniciating database
	// Initializes the database connection.
	log.Info("setting up database connection and running migrations")
	db, err := database.NewDatabaseWithMigrations(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Inicie o servidor
	e := echo.New()
	api.RegisterHTTPRoutes(e, db)
	e.Logger.Fatal(e.Start(":5802"))

}
