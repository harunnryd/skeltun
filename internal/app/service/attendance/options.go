package attendance

import (
	"skeltun/config"
	"skeltun/internal/app/repository"
)

// Option ...
type Option func(*Attendance)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(attendance *Attendance) {
		attendance.config = config
	}
}

// WithRepo ...
func WithRepo(repo repository.IRepository) Option {
	return func(attendance *Attendance) {
		attendance.repo = repo
	}
}
