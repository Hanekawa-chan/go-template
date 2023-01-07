package database

import (
	"github.com/rs/zerolog"
	"project-name/internal/app"
)

type adapter struct {
}

func NewAdapter(logger *zerolog.Logger, config *app.Config) (app.Database, error) {
	return &adapter{}, nil
}
