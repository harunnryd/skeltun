// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

import (
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"time"

	// main package.
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlOption ...
type MysqlOption struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	AddParams       string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

// OpenMysql ...
func OpenMysql(cfg config.IConfig) (dbx *gorm.DB, err error) {
	opt := mysqlConfigurer(cfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", opt.Username, opt.Password, opt.Host, opt.Port, opt.DBName, opt.AddParams)
	dbx, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	var sqlDB, _ = dbx.DB()

	sqlDB.SetConnMaxLifetime(time.Duration(opt.ConnMaxLifetime) * time.Second)
	sqlDB.SetMaxIdleConns(opt.MaxIdleConns)
	sqlDB.SetMaxOpenConns(opt.MaxOpenConns)
	return
}

func mysqlConfigurer(cfg config.IConfig) MysqlOption {
	return MysqlOption{
		Host:            cfg.GetString("database.mysql.host"),
		Port:            cfg.GetInt("database.mysql.port"),
		Username:        cfg.GetString("database.mysql.username"),
		Password:        cfg.GetString("database.mysql.password"),
		DBName:          cfg.GetString("database.mysql.db_name"),
		AddParams:       cfg.GetString("database.mysql.add_params"),
		MaxOpenConns:    cfg.GetInt("database.mysql.max_open_conns"),
		MaxIdleConns:    cfg.GetInt("database.mysql.max_idle_conns"),
		ConnMaxLifetime: cfg.GetInt("database.mysql.conn_max_lifetime"),
	}
}
