package configs

import (
	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		Http
		Postgres
		Logger
	}

	Http struct {
		Port string
	}
	Postgres struct {
		PostgresUrl     string
		PostgresMaxPool int
	}

	Logger struct {
		Level string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
