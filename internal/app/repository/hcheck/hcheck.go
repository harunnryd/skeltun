package hcheck

import (
	"fmt"
	"skeltun/config"
	"skeltun/internal/app/driver/db"

	"github.com/jmoiron/sqlx"
)

// IHcheck ...
type IHcheck interface {
	Ping(string) error
}

// Hcheck ...
type Hcheck struct {
	config  config.IConfig
	dbMysql *sqlx.DB
	dbPgsql *sqlx.DB
}

// New ...
func New(opts ...Option) IHcheck {
	hcheck := new(Hcheck)
	for _, opt := range opts {
		opt(hcheck)
	}
	return hcheck
}

// Ping ...
func (hcheck *Hcheck) Ping(dialect string) (err error) {
	if dialect == db.MysqlDialectParam {
		err = hcheck.dbMysql.Ping()
	}

	if err != nil {
		return
	}

	if dialect == db.PgsqlDialectParam {
		err = hcheck.dbPgsql.Ping()
	}

	fmt.Println("Print from hcheck repository ...")
	return
}
