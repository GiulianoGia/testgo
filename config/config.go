package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type MariaDBConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     int
}

type ServerConfig struct {
	DatabaseConnectionDetails MariaDBConfig
}

func NewServerConfig(ctx context.Context) *ServerConfig {
	config := &ServerConfig{
		DatabaseConnectionDetails: MariaDBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     3306,
		},
	}
	return config
}

func (c MariaDBConfig) DSN() string {

	cfg := mysql.NewConfig()

	cfg.DBName = c.Database

	cfg.ParseTime = true

	cfg.User = c.Username

	cfg.Passwd = c.Password

	cfg.Net = "tcp"

	cfg.Addr = fmt.Sprintf("%v:%v", c.Host, c.Port)

	return cfg.FormatDSN()

}
