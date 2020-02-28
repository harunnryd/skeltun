package service

import (
	"skeltun/internal/app/service/attendance"
	"skeltun/internal/app/service/hcheck"
)

// IService ...
type IService interface {
	Hcheck() hcheck.IHcheck
	Attendance() attendance.IAttendance
}

// Service ...
type Service struct {
	hcheck     hcheck.IHcheck
	attendance attendance.IAttendance
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

// Attendance ...
func (svc *Service) Attendance() attendance.IAttendance {
	return svc.attendance
}
