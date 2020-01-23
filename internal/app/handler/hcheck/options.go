package hcheck

import (
	"skeltun/config"
	"skeltun/internal/app/service"
)

// Option ...
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithService ...
func WithService(service service.IService) Option {
	return func(hcheck *Hcheck) {
		hcheck.service = service
	}
}
