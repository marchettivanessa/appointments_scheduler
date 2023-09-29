package main

import (
	"appointments_scheduler/api"
	"appointments_scheduler/config"
	"appointments_scheduler/database"
	"appointments_scheduler/logging"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	//Start config
	envPath := config.GetEnvOrDefault("ENV_FILE_PATH", ".env")
	c := config.MustParseConfig(envPath, "development")
	// Initializes logging.
	log.Info("setting up logging")
	logging.InitLogging(c.Log)

	// Iniciating database
	// Initializes the database connection.
	db, err := database.NewDatabaseWithMigrations(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Inicie o servidor
	e := echo.New()
	api.RegisterHTTPRoutes(e, db)
	e.Logger.Fatal(e.Start(":5802"))

}
