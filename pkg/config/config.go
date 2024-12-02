package config

import "github.com/spf13/viper"

// Config represents the environment variables.
type Config struct {
	APIPORT   string `mapstructure:"APIPORT"`
	SECRETKEY string `mapstructure:"JWTKEY"`
	//ADMINSECERETKEY string `mapstructure:"ADMINJWTKEY"`
	USERPORT    string `mapstructure:"USERPORT"`
	ADMINPORT   string `mapstructure:"ADMINPORT"`
	CHATPORT    string `mapstructure:"CHATPORT"`
	JWTISSUER   string `mapstructure:"JWT_ISSUER"`
	JWTAUDIENCE string `mapstructure:"JWT_AUDIENCE"`
	STRIPEKEY   string `maostructuee:"STRIPE_SECRET_KEY"`
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
