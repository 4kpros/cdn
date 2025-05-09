package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	// Application config
	AppPort  int    `mapstructure:"APP_PORT"`
	Hostname string `mapstructure:"HOST_NAME"`

	// API config
	ApiKey       string `mapstructure:"API_KEY"`
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`
}

var Env = &Environment{}

// LoadEnv Loads environment variables.
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
