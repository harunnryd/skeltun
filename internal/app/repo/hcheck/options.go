// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"

	"gorm.io/gorm"
)

// Option is a closure that is used for accessing the local variables.
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithDatabase ...
func WithDatabase(dialect string, conn *gorm.DB) Option {
	return func(hcheck *Hcheck) {
		if dialect == db.MysqlDialectParam {
			hcheck.ormMySQL = conn
		}
		if dialect == db.PgsqlDialectParam {
			hcheck.ormPgSQL = conn
		}
	}
}
