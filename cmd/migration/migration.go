// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"database/sql"
	"fmt"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"os"
	"time"

	// import src driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// IMigration ...
type IMigration interface {
	Up(string) error
	Down(string) error
	Create(string, string) error
}

// Migration ...
type Migration struct {
	ormPgSQL *gorm.DB
	ormMySQL *gorm.DB
	ormRaw   *sql.DB
}

// New ...
func New(opts ...Option) IMigration {
	migration := new(Migration)
	for _, opt := range opts {
		opt(migration)
	}
	return migration
}

// Up ...
func (migration *Migration) Up(dialect string) (err error) {
	mgr, err := migration.migrateInstance(dialect)
	if err != nil {
		return
	}
	defer mgr.Close()
	return mgr.Up()
}

// Down ...
func (migration *Migration) Down(dialect string) (err error) {
	mgr, err := migration.migrateInstance(dialect)
	if err != nil {
		return
	}
	defer mgr.Close()
	return mgr.Down()
}

// Create ...
func (migration *Migration) Create(name string, ext string) (err error) {
	base := fmt.Sprintf("%v/%v_%v.", "migration/sql", time.Now().Unix(), name)
	err = migration.createFile(base + "up." + ext)
	if err != nil {
		return
	}

	err = migration.createFile(base + "down." + ext)
	return
}

func (migration *Migration) createFile(fname string) (err error) {
	_, err = os.Create(fname)
	return
}

func (migration *Migration) migrateInstance(dialect string) (mgr *migrate.Migrate, err error) {
	driver, err := migration.getDriver(dialect)
	if err != nil {
		return
	}

	mgr, err = migrate.NewWithDatabaseInstance(
		migration.getSource("migration/sql"),
		dialect,
		driver,
	)
	if err != nil {
		return
	}

	return
}

func (migration *Migration) getDriver(dialect string) (driver database.Driver, err error) {
	if dialect == db.PgsqlDialectParam {
		migration.ormRaw, _ = migration.ormPgSQL.DB()
		driver, err = postgres.WithInstance(migration.ormRaw, &postgres.Config{})
	}
	if dialect == db.MysqlDialectParam {
		migration.ormRaw, _ = migration.ormMySQL.DB()
		driver, err = mysql.WithInstance(migration.ormRaw, &mysql.Config{})
	}
	return
}

func (migration *Migration) getSource(src string) string {
	return fmt.Sprintf("file://%s", src)
}
