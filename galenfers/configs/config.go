package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl     string
	JWTSecret string
}

func Load() Config {
	viper.SetConfigFile("galenfers/configs/.env")
	viper.ReadInConfig()
	DB_URL := viper.Get("DB_URL").(string)
	JWT_SECRET := viper.Get("JWT_SECRET").(string)

	return Config{
		DBUrl:     DB_URL,
		JWTSecret: JWT_SECRET,
	}
}
