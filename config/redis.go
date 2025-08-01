package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RedisEnv struct {
	Host string
	Port string
}

var redisConfig RedisEnv

func LoadRedisEnv() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	redisConfig = RedisEnv{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
	}
}

func GetRedisEnv() RedisEnv {
	return redisConfig
}
