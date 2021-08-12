package utils

import (
	"os"
)

// APIConfig ...
type APIConfig struct {
	HeaderValue struct {
		ClientID string
	}
}

// ConfigDB ...
type ConfigDB struct {
	MySQL struct {
		DBMS     string
		Protocol string
		Username string
		Password string
		DBName   string
	}
}

// NewAPIConfig ...
func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	config.HeaderValue.ClientID = os.Getenv("SBCNTR_CLIENT_ID_HEADER")

	return config
}

// NewConfigDB ...
func NewConfigDB() *ConfigDB {
	config := new(ConfigDB)

	//os.Setenv("DB_HOST", "localhost")
	//os.Setenv("DB_USERNAME", "sbcntruser")
	//os.Setenv("DB_PASSWORD", "sbcntrpass")
	//os.Setenv("DB_NAME", "sbcntrapp")

	config.MySQL.DBMS = "mysql"
	config.MySQL.Protocol = "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	config.MySQL.Username = os.Getenv("DB_USERNAME")
	config.MySQL.Password = os.Getenv("DB_PASSWORD")
	config.MySQL.DBName = os.Getenv("DB_NAME")

	return config
}
