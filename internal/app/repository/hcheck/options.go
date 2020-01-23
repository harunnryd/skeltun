package hcheck

import (
	"skeltun/config"
	"skeltun/internal/app/driver/db"

	"github.com/jmoiron/sqlx"
)

// Option ...
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithDatabase ...
func WithDatabase(dialect string, conn *sqlx.DB) Option {
	return func(hcheck *Hcheck) {
		if dialect == db.MysqlDialectParam {
			hcheck.dbMysql = conn
		}
		if dialect == db.PgsqlDialectParam {
			hcheck.dbPgsql = conn
		}
	}
}
