package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	config := &Config{
		AppPort:    viper.GetString("APP_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetInt("DB_PORT"),
	}

	return config, nil
}
