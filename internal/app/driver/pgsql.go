// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

import (
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"time"

	"gorm.io/gorm"

	// postgres
	"gorm.io/driver/postgres"
)

// PgsqlOption ...
type PgsqlOption struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	MaxPoolSize time.Duration
	Sslmode     string
}

// OpenPgsql ...
func OpenPgsql(cfg config.IConfig) (dbx *gorm.DB, err error) {
	opt := pgsqlConfigurer(cfg)
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", opt.Host, opt.Port, opt.Username, opt.DBName, opt.Password, opt.Sslmode)
	dbx, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	var sqlDB, _ = dbx.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opt.MaxPoolSize)
	return
}

func pgsqlConfigurer(cfg config.IConfig) PgsqlOption {
	return PgsqlOption{
		Host:        cfg.GetString("database.pgsql.host"),
		Port:        cfg.GetInt("database.pgsql.port"),
		Username:    cfg.GetString("database.pgsql.username"),
		Password:    cfg.GetString("database.pgsql.password"),
		DBName:      cfg.GetString("database.pgsql.db_name"),
		MaxPoolSize: cfg.GetDuration("database.pgsql.max_pool_size"),
		Sslmode:     cfg.GetString("database.pgsql.sslmode"),
	}
}
