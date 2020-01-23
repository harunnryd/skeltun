package db

import (
	"errors"
	"skeltun/config"
	"skeltun/internal/app/driver"

	"github.com/jmoiron/sqlx"
)

// IDB ...
type IDB interface {
	Manager(string) (*sqlx.DB, error)
}

// DB ...
type DB struct {
	config config.IConfig
}

// New ...
func New(callbacks ...Option) IDB {
	db := new(DB)
	for _, callback := range callbacks {
		callback(db)
	}
	return db
}

// Manager ...
func (db *DB) Manager(dialect string) (dbraw *sqlx.DB, err error) {
	switch dialect {
	case MysqlDialectParam:
		dbraw, err = driver.OpenMysql(db.config)
	case PgsqlDialectParam:
		dbraw, err = driver.OpenPgsql(db.config)
	default:
		err = errors.New("Undefined connection; ignore error if desired")
	}
	return
}
