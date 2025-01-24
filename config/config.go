package config

import "github.com/spf13/viper"

type Config struct {
	Database_url string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	// var config Config
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()
	// viper.Unmarshal(&config)
	return &config
}
