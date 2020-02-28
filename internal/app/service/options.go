package service

import (
	"skeltun/config"
	"skeltun/internal/app/repository"
	"skeltun/internal/app/service/attendance"
	"skeltun/internal/app/service/hcheck"
)

// Option ...
type Option func(*Service)

// WithService ...
func WithService(config config.IConfig) Option {
	repo := repository.New(repository.WithDatabase(config))

	return func(svc *Service) {
		// Inject all your service's in here.
		// Example :
		// svc.user = user.New(
		//    user.WithConfig(config),
		//    user.WithRepo(repo),
		// )
		svc.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithRepo(repo),
		)
		svc.attendance = attendance.New(
			attendance.WithConfig(config),
			attendance.WithRepo(repo),
		)
	}
}
