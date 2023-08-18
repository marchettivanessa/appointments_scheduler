package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Env      string
	AppPort  string
	Database DatabaseConfig
	Log      LogConfig
}

type DatabaseConfig struct {
	Name           string
	Host           string
	Password       string
	Username       string
	MigrationPath  string
	Port           int
	ConnectTimeout int
	Schema         string
}

type LogConfig struct {
	AppName            string
	Level              string
	EnableSyslog       bool
	EnableReportCaller bool
}

func MustParseConfig() Config {
	return Config{
		Env:     MustGetEnv("ENV"),
		AppPort: MustGetEnv("APP_PORT"),
		Database: DatabaseConfig{
			Name:           MustGetEnv("DATABASE_NAME"),
			Host:           MustGetEnv("DATABASE_HOST"),
			Password:       MustGetEnv("DATABASE_PASSWORD"),
			Username:       MustGetEnv("DATABASE_USERNAME"),
			MigrationPath:  MustGetEnv("DATABASE_MIGRATION_PATH"),
			Port:           MustParseInt("DATABASE_PORT"),
			ConnectTimeout: MustParseInt("DATABASE_CONNECT_TIMEOUT"),
			Schema:         MustGetEnv("DATABASE_SCHEMA"),
		},
		Log: LogConfig{
			AppName:            MustGetEnv("APP_NAME"),
			Level:              MustGetEnv("LOG_LEVEL"),
			EnableSyslog:       MustParseBool(MustGetEnv("ENABLE_SYSLOG")),
			EnableReportCaller: MustParseBool(MustGetEnv("ENABLE_REPORT_CALLER")),
		},
	}
}

func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.WithField("key", key).Fatal("failed to find environment variable")
	}
	return v
}

func MustParseBool(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.WithError(err).WithField("value", v).Fatal("failed converting value to bool")
	}
	return b
}

func MustParseInt(key string) int {
	v, err := strconv.Atoi(MustGetEnv(key))
	if err != nil {
		log.WithError(err).Fatal("no valid value assigned to env variable", key)
	}
	return v
}
