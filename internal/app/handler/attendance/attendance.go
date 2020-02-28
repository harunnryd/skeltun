package attendance

import (
	"net/http"
	"skeltun/config"
	"skeltun/internal/app/service"
)

// IAttendance ...
type IAttendance interface {
	AttendanceBook(http.ResponseWriter, *http.Request) (interface{}, error)
}

// Attendance ...
type Attendance struct {
	config  config.IConfig
	service service.IService
}

// New ...
func New(opts ...Option) IAttendance {
	attendance := new(Attendance)
	for _, opt := range opts {
		opt(attendance)
	}
	return attendance
}

// AttendanceBook ...
func (attendance *Attendance) AttendanceBook(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	data, err = attendance.service.Attendance().GetAttendance(r.Context(), r.URL.Query().Get("date"))
	return
}
