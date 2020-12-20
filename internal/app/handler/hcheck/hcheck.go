// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"net/http"
)

// IHcheck ...
type IHcheck interface {
	HealthCheck(http.ResponseWriter, *http.Request) (interface{}, error)
}

// Hcheck ...
type Hcheck struct {
	config  config.IConfig
	usecase usecase.IUseCase
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
		err = hcheck.usecase.GetHcheck().PgsqlDB()
	}

	if err != nil {
		return
	}

	if hcheck.config.GetBool("database.mysql.is_active") {
		err = hcheck.usecase.GetHcheck().MysqlDB()
	}

	w.Header().Add("Content-Type", "application/json")

	return
}
