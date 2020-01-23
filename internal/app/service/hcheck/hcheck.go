package hcheck

import (
	"fmt"
	"skeltun/config"
	"skeltun/internal/app/driver/db"
	"skeltun/internal/app/repository"
)

// IHcheck ...
type IHcheck interface {
	MysqlDB() error
	PgsqlDB() error
}

// Hcheck ...
type Hcheck struct {
	config  config.IConfig
	repo   repository.IRepository
}

// New ...
func New(opts ...Option) IHcheck {
	hcheck := new(Hcheck)
	for _, opt := range opts {
		opt(hcheck)
	}
	return hcheck
}

// MysqlDB ...
func (hcheck *Hcheck) MysqlDB() error {
	return hcheck.repo.Hcheck().Ping(db.MysqlDialectParam)
}

// PgsqlDB ...
func (hcheck *Hcheck) PgsqlDB() error {
	fmt.Println("Print from hcheck service ...")
	return hcheck.repo.Hcheck().Ping(db.PgsqlDialectParam)
}
