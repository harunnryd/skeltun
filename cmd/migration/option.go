package migration

import (
	"skeltun/internal/app/driver/db"

	"github.com/jmoiron/sqlx"
)

// Option ...
type Option func(*Migration)

// WithDatabase ...
func WithDatabase(dialect string, conn *sqlx.DB) Option {
	return func(migration *Migration) {
		if dialect == db.PgsqlDialectParam {
			migration.dbPgsql = conn
		}
		if dialect == db.MysqlDialectParam {
			migration.dbMysql = conn
		}
	}
}
