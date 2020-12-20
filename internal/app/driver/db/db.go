// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"errors"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver"

	"gorm.io/gorm"
)

// IDB ...
type IDB interface {
	Manager(string) (*gorm.DB, error)
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
func (db *DB) Manager(dialect string) (dbraw *gorm.DB, err error) {
	switch dialect {
	case MysqlDialectParam:
		if !db.config.GetBool("database.mysql.is_active") {
			return
		}
		dbraw, err = driver.OpenMysql(db.config)
	case PgsqlDialectParam:
		if !db.config.GetBool("database.pgsql.is_active") {
			return
		}
		dbraw, err = driver.OpenPgsql(db.config)
	default:
		err = errors.New("Undefined connection; ignore error if desired")
	}

	return
}
