// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"

	"github.com/gocraft/work"
)

// IHcheck ...
type IHcheck interface {
	MysqlDB() error
	PgsqlDB() error
}

// Hcheck ...
type Hcheck struct {
	config config.IConfig
	repo   repo.IRepo
	pkg    pkg.IPkg
	job    job.IJob
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
	return hcheck.repo.GetHcheck().Ping(db.MysqlDialectParam)
}

// PgsqlDB ...
func (hcheck *Hcheck) PgsqlDB() error {
	fmt.Println("Print from hcheck usecase ...")
	hcheck.job.Dispatch("hcheck", 10, work.Q{"response_code": "000000", "response_desc": "Success"})
	return hcheck.repo.GetHcheck().Ping(db.PgsqlDialectParam)
}
