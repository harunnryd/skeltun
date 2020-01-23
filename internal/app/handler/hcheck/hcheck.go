package hcheck

import (
	"net/http"
	"skeltun/config"
	"skeltun/internal/app/service"
)

// IHcheck ...
type IHcheck interface {
	HealthCheck(http.ResponseWriter, *http.Request) (interface{}, error)
}

// Hcheck ...
type Hcheck struct {
	config  config.IConfig
	service service.IService
}

// New ...
func New(opts ...Option) IHcheck {
	hcheck := new(Hcheck)
	for _, opt := range opts {
		opt(hcheck)
	}
	return hcheck
}

// HealthCheck ...
func (hcheck *Hcheck) HealthCheck(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	if hcheck.config.GetBool("database.pgsql.is_active") {
		err = hcheck.service.Hcheck().PgsqlDB()
	}

	if err != nil {
		return
	}

	if hcheck.config.GetBool("database.mysql.is_active") {
		err = hcheck.service.Hcheck().MysqlDB()
	}

	return
}
