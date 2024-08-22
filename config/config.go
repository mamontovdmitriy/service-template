package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	Log  `yaml:"log"`
	PG   `yaml:"postgres"`
}

type App struct {
	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
}

type PG struct {
	MaxPoolSize int    `env-required:"true" yaml:"max_pool_size" env:"PG_MAX_POOL_SIZE"`
	URL         string `env-required:"true" env:"PG_URL"`
}

func NewConfig(pathConfig string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", pathConfig), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
