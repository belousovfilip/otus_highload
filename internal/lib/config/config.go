package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTPServer struct {
			Addr string
		}
		DBPrimary struct {
			Connection string
		}
	}
)

func NewConfig() *Config {
	return &Config{}
}

func InitAppConfigFromEnv() (*Config, error) {
	options := struct {
		DBPrimaryConnection string `env:"DB_PRIMARY_CONNECTION"`
		HTTPServerAddr      string `env:"HTTP_SERVER_ADDR" env-default:":8083"`
	}{}
	if err := cleanenv.ReadEnv(&options); err != nil {
		return nil, errors.New("cannot read config: %s" + err.Error())
	}
	config := NewConfig()
	config.DBPrimary.Connection = options.DBPrimaryConnection
	config.HTTPServer.Addr = options.HTTPServerAddr
	return config, nil
}
