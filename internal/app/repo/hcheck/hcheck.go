// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"

	"gorm.io/gorm"
)

// IHcheck ...
type IHcheck interface {
	Ping(string) error
}

// Hcheck ...
type Hcheck struct {
	config   config.IConfig
	ormMySQL *gorm.DB
	ormPgSQL *gorm.DB
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
		var sqlDB, _ = hcheck.ormMySQL.DB()
		if err = sqlDB.Ping(); err != nil {
			return
		}
	}

	if dialect == db.PgsqlDialectParam {
		var sqlDB, _ = hcheck.ormPgSQL.DB()
		if err = sqlDB.Ping(); err != nil {
			return
		}
	}

	fmt.Println("Print from hcheck repo ...")
	return
}
