package config

import "github.com/spf13/viper"

// Config represents the environment variables.
type Config struct {
	APIPORT   string `mapstructure:"APIPORT"`
	SECRETKEY string `mapstructure:"JWTKEY"`
	USERPORT  string `mapstructure:"USERPORT"`
	ADMINPORT string `mapstructure:"ADMINPORT"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
