package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseDriver string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost   string `mapstructure:"DATABASE_HOST"`
	DatabasePort   uint   `mapstructure:"DATABASE_PORT"`
	DatabaseUser   string `mapstructure:"DATABASE_USER"`
	DatabasePass   string `mapstructure:"DATABASE_PASS"`
	DatabaseName   string `mapstructure:"DATABASE_NAME"`
}

var cfg *Config

func LoadConfig(path string) *Config {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
