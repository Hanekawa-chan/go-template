package app

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"project-name/internal/services/config"
)

type Config struct {
	Logger     *LoggerConfig
	HTTPServer *config.HTTPConfig
}

type LoggerConfig struct {
	LogLevel string `default:"debug"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	logger := LoggerConfig{}

	err := envconfig.Process("kanji_auth", &logger)
	if err != nil {
		log.Err(err).Msg("logger config error")
		return nil, err
	}

	cfg.Logger = &logger

	return &cfg, nil
}
