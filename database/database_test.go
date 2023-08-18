package database

import (
	"appointments_scheduler/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	t.Setenv("DATABASE_NAME", "database_name")
	t.Setenv("DATABASE_HOST", "localhost")
	t.Setenv("DATABASE_PASSWORD", "database_password")
	t.Setenv("DATABASE_USERNAME", "database_username")
	t.Setenv("DATABASE_MIGRATION_PATH", "database_migration_path")
	t.Setenv("DATABASE_PORT", "1338")
	t.Setenv("DATABASE_CONNECT_TIMEOUT", "10")
	t.Setenv("DATABASE_SCHEMA", "schema")

	databaseConfig := config.DatabaseConfig{
		Name:           "database_name",
		Host:           "localhost",
		Password:       "database_password",
		Username:       "database_username",
		MigrationPath:  "database_migration_path",
		Port:           1338,
		ConnectTimeout: 10,
		Schema:         "schema",
	}

	expectedConnectionString := "host=localhost port=1338 username=database_username password=password dbname=database_name connect_timeout=10 search_path=public sslmode=disable"
	//	connectionString, _ := NewDabataseWithMigrations(databaseConfig)

	dbc, _ := newDatabaseConn(databaseConfig)
	assert.Equal(t, expectedConnectionString, dbc)
}
