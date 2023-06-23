package database

import (
	"appointments_scheduler/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Postgres = "postgres"
)

type Database struct {
	Config     config.DatabaseConfig
	Connection *sqlx.DB
}

func newDatabase(dbc config.DatabaseConfig) (*Database, error) {
	connectionString := buildConnectionString(dbc)
	databaseConnection, err := sqlx.Open(Postgres, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to create the database connection: %w", err)
	}

	return &Database{
		Config:     dbc,
		Connection: databaseConnection,
	}, nil
}

func buildConnectionString(dbc config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d search_path=public,%s sslmode=disable",
		dbc.Host,
		dbc.Port,
		dbc.Username,
		dbc.Password,
		dbc.Name,
		dbc.ConnectTimeout,
		dbc.Schema,
	)
}
