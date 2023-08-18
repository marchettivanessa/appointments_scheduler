package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParseConfig(t *testing.T) {
	t.Setenv("DATABASE_NAME", "database_name")
	t.Setenv("DATABASE_HOST", "localhost")
	t.Setenv("DATABASE_PASSWORD", "database_password")
	t.Setenv("DATABASE_USERNAME", "database_username")
	t.Setenv("DATABASE_MIGRATION_PATH", "database_migration_path")
	t.Setenv("DATABASE_PORT", "1338")
	t.Setenv("DATABASE_CONNECT_TIMEOUT", "10")
	t.Setenv("DATABASE_SCHEMA", "schema")
	t.Setenv("APP_NAME", "app_name")
	t.Setenv("LOG_LEVEL", "log_level")
	t.Setenv("ENABLE_REPORT_CALLER", "false")
	t.Setenv("ENABLE_SYSLOG", "false")
	t.Setenv("ENV", "development")
	t.Setenv("APP_PORT", "81")
	databaseConfig := DatabaseConfig{
		Name:           "database_name",
		Host:           "localhost",
		Password:       "database_password",
		Username:       "database_username",
		MigrationPath:  "database_migration_path",
		Port:           1338,
		ConnectTimeout: 10,
		Schema:         "schema",
	}

	logConfig := LogConfig{
		AppName:            "app_name",
		Level:              "log_level",
		EnableSyslog:       false,
		EnableReportCaller: false,
	}
	expectedConfig := Config{
		Env:      "development",
		AppPort:  "81",
		Database: databaseConfig,
		Log:      logConfig,
	}

	config := MustParseConfig()
	assert.Equal(t, expectedConfig, config)
}
