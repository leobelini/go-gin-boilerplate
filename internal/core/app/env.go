package app

import (
	"leobelini/cashly/internal/core/dto"

	"github.com/spf13/viper"
)

func LoadEnv() (*dto.DtoEnvApp, error) {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	var AppEnv = dto.DtoEnvApp{
		Server: dto.DtoEnvAppServer{
			Port: viper.GetInt("server.port"),
			Host: viper.GetString("server.host"),
		},
		Database: dto.DtoEnvAppDatabase{
			File:        viper.GetString("database.file"),
			AutoMigrate: viper.GetBool("database.autoMigrate"),
		},
		Redis: dto.DtoEnvAppRedis{
			Host: viper.GetString("redis.host"),
			Port: viper.GetInt("redis.port"),
		},
		IsProd: viper.GetBool("isProduction"),
		App: dto.DtoEnvAppApp{
			Name: viper.GetString("app.name"),
			URL:  viper.GetString("app.url"),
		},
		Smtp: dto.DtoEnvAppSmtp{
			Host: viper.GetString("smtp.host"),
			Port: viper.GetInt("smtp.port"),
			From: viper.GetString("smtp.from"),
		},
	}

	return &AppEnv, nil
}
