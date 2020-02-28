package handler

import (
	"skeltun/internal/app/handler/attendance"
	"skeltun/internal/app/handler/hcheck"
)

// IHandler ...
type IHandler interface {
	Hcheck() hcheck.IHcheck
	Attendance() attendance.IAttendance
}

// Handler ...
type Handler struct {
	hcheck     hcheck.IHcheck
	attendance attendance.IAttendance
}

// New ...
func New(opts ...Option) IHandler {
	handler := new(Handler)
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// Hcheck ...
func (handler *Handler) Hcheck() hcheck.IHcheck {
	return handler.hcheck
}

// Attendance ...
func (handler *Handler) Attendance() attendance.IAttendance {
	return handler.attendance
}
