package attendance

import (
	"skeltun/config"
	"skeltun/internal/app/driver/db"

	"github.com/jmoiron/sqlx"
)

// Option ...
type Option func(*Attendance)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(attendance *Attendance) {
		attendance.config = config
	}
}

// WithDatabase ...
func WithDatabase(dialect string, conn *sqlx.DB) Option {
	return func(attendance *Attendance) {
		if dialect == db.MysqlDialectParam {
			attendance.dbMysql = conn
		}
		if dialect == db.PgsqlDialectParam {
			attendance.dbPgsql = conn
		}
	}
}
