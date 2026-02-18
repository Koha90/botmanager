// Package config represantation a configuration of application.
package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string `env:"ENV"  env-default:"local"`
	HTTP struct {
		Port string `env:"ENV_PORT" env-default:"8080"`
	} `env:"HTTP"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Errorf("cannot read config: %w", err))
	}

	return &cfg
}
