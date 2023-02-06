package httpserver

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"go.uber.org/ratelimit"
	"net/http"
	"project-name/internal/app"
)

type adapter struct {
	logger  *zerolog.Logger
	config  *Config
	server  *http.Server
	service app.Service
	limiter ratelimit.Limiter
}

func NewAdapter(logger *zerolog.Logger, config *Config, service app.Service) app.HTTPServer {
	a := &adapter{
		logger:  logger,
		config:  config,
		service: service,
		limiter: ratelimit.New(config.RateLimit),
	}

	r := chi.NewRouter()

	a.initMiddlewares(r)

	r.Handle("/metrics", promhttp.Handler())
	r.MethodFunc(http.MethodGet, "/health-check", a.HealthCheck)

	a.server = &http.Server{
		Addr:    config.Address,
		Handler: r,
	}

	return a
}

func (a *adapter) ListenAndServe() error {
	a.logger.Info().Msgf("Listening and serving HTTP requests on: %v", a.config.Address)

	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.logger.Error().Err(err).Msg("Error listening and serving HTTP requests.")
		return err
	}

	return nil
}

func (a *adapter) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error().Err(err).Msg("Error shutting down HTTP adapter!")
		return err
	}

	return nil
}

func (a *adapter) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
