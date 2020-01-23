package handler

import (
	"skeltun/config"
	"skeltun/internal/app/handler/hcheck"
	"skeltun/internal/app/service"
)

// Option ...
type Option func(*Handler)

// WithHandler ...
func WithHandler(config config.IConfig) Option {
	service := service.New(service.WithService(config))
	return func(handler *Handler) {
		// Inject all your handler's in here.
		// Example :
		// handler.user = user.New(
		//     user.WithConfig(config),
		//     user.WithService(service),
		// )
		handler.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithService(service),
		)
	}
}
