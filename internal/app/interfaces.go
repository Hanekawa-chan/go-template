package app

import "context"

type Service interface {
}

type Database interface {
}

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}
