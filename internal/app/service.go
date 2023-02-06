package app

import (
	"github.com/rs/zerolog"
	"project-name/internal/app/config"
)

type service struct {
}

func NewService(logger *zerolog.Logger, config *config.Config) Service {
	return &service{}
}
