package repository

import (
	"skeltun/internal/app/repository/attendance"
	"skeltun/internal/app/repository/hcheck"
)

// IRepository ...
type IRepository interface {
	Hcheck() hcheck.IHcheck
	Attendance() attendance.IAttendance
}

// Repository ...
type Repository struct {
	hcheck     hcheck.IHcheck
	attendance attendance.IAttendance
}

// New ...
func New(opts ...Option) IRepository {
	repo := new(Repository)
	for _, opt := range opts {
		opt(repo)
	}
	return repo
}

// Hcheck ...
func (repo *Repository) Hcheck() hcheck.IHcheck {
	return repo.hcheck
}

// Attendance ...
func (repo *Repository) Attendance() attendance.IAttendance {
	return repo.attendance
}
