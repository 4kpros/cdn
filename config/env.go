package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	// Application config
	AppPort int    `mapstructure:"APP_PORT"`
	AppName string `mapstructure:"APP_NAME"`

	// API config
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`
}

var Env = &Environment{}

// Loads environment variables.
func LoadEnv() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(Env)
	}
	return err
}
