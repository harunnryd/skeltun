package migration

import (
	"fmt"
	"os"
	"skeltun/internal/app/driver/db"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/jmoiron/sqlx"
)

// IMigration ...
type IMigration interface {
	Up(string) error
	Down(string) error
	Create(string, string) error
}

// Migration ...
type Migration struct {
	dbPgsql *sqlx.DB
	dbMysql *sqlx.DB
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
		driver, err = postgres.WithInstance(migration.dbPgsql.DB, &postgres.Config{})
	}
	if dialect == db.MysqlDialectParam {
		driver, err = mysql.WithInstance(migration.dbMysql.DB, &mysql.Config{})
	}
	return
}

func (migration *Migration) getSource(src string) string {
	return fmt.Sprintf("file://%s", src)
}
