package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	AppName         string `env:"APP_NAME" envDefault:"blueprint-roa-golang"`
	Port            int    `env:"PORT" envDefault:"3000"`
	HttpPort        int    `env:"HTTP_PORT" envDefault:"3001"`
	Env             string `env:"ENV" envDefault:"localhost"`
	DbDriver        string `env:"DB_DRIVER" envDefault:"postgres"`
	DbHost          string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort          string `env:"DB_PORT" envDefault:"5432"`
	DbUser          string `env:"DB_USER" envDefault:"postgres"`
	DbName          string `env:"DB_NAME" envDefault:"products"`
	DbPassword      string `env:"DB_PASSWORD" envDefault:"postgres"`
	JaegerAgentHost string `env:"JAEGER_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_PORT" envDefault:"6831"`
}

func NewConfiguration() Configuration {
	config := Configuration{}

	if err := env.Parse(&config); err != nil {
		fmt.Errorf("%+v\n", err)
	}

	return config
}
