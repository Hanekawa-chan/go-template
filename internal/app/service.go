package app

import "github.com/rs/zerolog"

type service struct {
}

func NewService(logger *zerolog.Logger, config *Config) Service {
	return &service{}
}
