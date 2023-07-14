package main

import (
	"appointments_scheduler/config"
	"appointments_scheduler/database"

	"github.com/labstack/gommon/log"
)

func main() {
c := config.MustParseConfig()
	//Start config

	// Iniciating database
	// Initializes the database connection.
	log.Info("setting up database connection and running migrations")
	db, err := database.NewDabataseWithMigrations(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	//Iniciate logs

}
