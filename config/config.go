package config

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
