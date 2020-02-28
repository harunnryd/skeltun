package attendance

import (
	"skeltun/config"
	"skeltun/internal/app/service"
)

// Option ...
type Option func(*Attendance)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(attendance *Attendance) {
		attendance.config = config
	}
}

// WithService ...
func WithService(service service.IService) Option {
	return func(attendance *Attendance) {
		attendance.service = service
	}
}
