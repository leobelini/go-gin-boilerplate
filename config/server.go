package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerEnv struct {
	Port int
	Host string
}

var config ServerEnv

func LoadServerEnv() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config = ServerEnv{
		Port: viper.GetInt("server.port"),
		Host: viper.GetString("server.host"),
	}
}

func GetServerEnv() ServerEnv {
	return config
}
