package service

import (
	"skeltun/internal/app/service/hcheck"
)

// IService ...
type IService interface {
	Hcheck() hcheck.IHcheck
}

// Service ...
type Service struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IService {
	svc := new(Service)
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

// Hcheck ...
func (svc *Service) Hcheck() hcheck.IHcheck {
	return svc.hcheck
}
