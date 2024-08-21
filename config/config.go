package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.AddConfigPath("config")
	viper.SetConfigName("config." + env)

	return viper.ReadInConfig()
}
