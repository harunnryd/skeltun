package attendance

import (
	"context"
	"skeltun/config"

	"github.com/jmoiron/sqlx"
)

// IAttendance ...
type IAttendance interface {
	GetAttendance(ctx context.Context, date string) (dest []SkeltunGetAttendance, err error)
}

// Attendance ...
type Attendance struct {
	config  config.IConfig
	dbMysql *sqlx.DB
	dbPgsql *sqlx.DB
	statement
}

type statement struct {
	getAttendance *sqlx.NamedStmt
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
func (attendance *Attendance) GetAttendance(ctx context.Context, date string) (dest []SkeltunGetAttendance, err error) {
	if attendance.statement.getAttendance == nil {
		attendance.statement.getAttendance, err = attendance.dbPgsql.PrepareNamedContext(ctx, skeltunGetAttendance)
		if err != nil {
			return
		}
	}

	args := map[string]interface{}{
		"date": date,
	}

	err = attendance.statement.getAttendance.SelectContext(ctx, &dest, args)

	return
}
