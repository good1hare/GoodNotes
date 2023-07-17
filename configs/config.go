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
		Telegram
	}

	Http struct {
		Port string
	}
	Postgres struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Url     string `env-required:"true"                 env:"PG_URL"`
	}

	Logger struct {
		Level string
	}

	Telegram struct {
		Token string
		Debug bool
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
