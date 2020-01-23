package hcheck

import (
	"skeltun/config"
	"skeltun/internal/app/repository"
)

// Option ...
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithRepo ...
func WithRepo(repo repository.IRepository) Option {
	return func(hcheck *Hcheck) {
		hcheck.repo = repo
	}
}
