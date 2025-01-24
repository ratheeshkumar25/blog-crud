package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database_url string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() *Config {
	var config Config

	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	return &config
}
