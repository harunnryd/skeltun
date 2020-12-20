// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/harunnryd/skeltun/internal/app/driver/db"

	"gorm.io/gorm"
)

// Option ...
type Option func(*Migration)

// WithDatabase ...
func WithDatabase(dialect string, conn *gorm.DB) Option {
	return func(migration *Migration) {
		if dialect == db.PgsqlDialectParam {
			migration.ormPgSQL = conn
		}
		if dialect == db.MysqlDialectParam {
			migration.ormMySQL = conn
		}
	}
}
