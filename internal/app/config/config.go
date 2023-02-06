package config

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
	project := "kanji_project_name"

	err := envconfig.Process(project, &logger)
	if err != nil {
		log.Err(err).Msg("logger config error")
		return nil, err
	}

	cfg.Logger = &logger

	return &cfg, nil
}
