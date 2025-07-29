package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseEnv struct {
	Filename    string
	AutoMigrate bool
}

var databaseConfig DatabaseEnv

func LoadDatabaseEnv() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	databaseConfig = DatabaseEnv{
		Filename:    viper.GetString("database.file"),
		AutoMigrate: viper.GetBool("database.autoMigrate"),
	}
}

func GetDatabaseEnv() DatabaseEnv {
	return databaseConfig
}
