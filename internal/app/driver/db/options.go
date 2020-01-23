package db

import "skeltun/config"

const (
	// MysqlDialectParam ...
	MysqlDialectParam = "mysql"
	// PgsqlDialectParam ...
	PgsqlDialectParam = "postgres"
)

// Option ...
type Option func(*DB)

// WithConfig ....
func WithConfig(config config.IConfig) Option {
	return func(db *DB) {
		db.config = config
	}
}
