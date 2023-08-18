package database

import (
	"appointments_scheduler/config"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

const (
	Postgres = "postgres"
)

type Database struct {
	Config     config.DatabaseConfig
	Connection *sqlx.DB
}

func newDatabaseConn(dbc config.DatabaseConfig) (*Database, error) {
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
	return fmt.Sprintf("host=%s port=%d username=%s password=%s dbname=%s connect_timeout=%d search_path=public,%s sslmode=disable",
		dbc.Host,
		dbc.Port,
		dbc.Username,
		dbc.Password,
		dbc.Name,
		dbc.ConnectTimeout,
		dbc.Schema,
	)
}

func (db *Database) Migrate() error {
	driver, err := postgres.WithInstance(db.Connection.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres drive: %w", err)
	}

	migrationPath := fmt.Sprintf("file://%s", db.Config.MigrationPath)
	migration, err := migrate.NewWithDatabaseInstance(migrationPath, Postgres, driver)
	if err != nil {
		return fmt.Errorf("failed initializing database migration: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed migrating the database: %w", err)
	}

	log.Info("Database migration successful")
	return nil
}

func NewDabataseWithMigrations(c config.DatabaseConfig) (*Database, error) {
	database, err := newDatabaseConn(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create the database connection: %w", err)
	}
	err = database.Migrate()
	if err != nil {
		return nil, fmt.Errorf("failed to migrate the database: %w", err)
	}
	return database, nil
}

func (db *Database) ResetMigration() error {
	updateStatement := "DROP TABLE IF EXISTS schema_migrations"
	_, err := db.Connection.Exec(updateStatement)
	if err != nil {
		return fmt.Errorf("failed reseting migrations: %w", err)
	}
	return nil
}

