package attendance

import (
	"context"
	"skeltun/config"
	"skeltun/internal/app/repository"
)

// IAttendance ...
type IAttendance interface {
	GetAttendance(ctx context.Context, date string) (dest interface{}, err error)
}

// Attendance ...
type Attendance struct {
	config config.IConfig
	repo   repository.IRepository
}

// New ...
func New(opts ...Option) IAttendance {
	attendance := new(Attendance)
	for _, opt := range opts {
		opt(attendance)
	}
	return attendance
}

// GetAttendance ...
func (attendance *Attendance) GetAttendance(ctx context.Context, date string) (dest interface{}, err error) {
	dest, err = attendance.repo.Attendance().GetAttendance(ctx, date)
	return
}
