package configs

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	DBDns       string `mapstructure:"DB_DNS"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path string) (*Config, error) {
	var config *Config

	viper.SetConfigName("configs")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic("Config file not found")
		}
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
