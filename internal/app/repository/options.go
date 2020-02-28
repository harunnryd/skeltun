package repository

import (
	"skeltun/config"
	"skeltun/internal/app/driver/db"
	"skeltun/internal/app/repository/attendance"
	"skeltun/internal/app/repository/hcheck"
)

// Option ...
type Option func(*Repository)

// WithDatabase ...
func WithDatabase(config config.IConfig) Option {
	dbase := db.New(db.WithConfig(config))
	mysqlConn, _ := dbase.Manager(db.MysqlDialectParam)
	pgsqlConn, _ := dbase.Manager(db.PgsqlDialectParam)

	return func(repo *Repository) {
		// Inject all your repo's in here.
		// Example :
		// repo.cache = cache.New(
		//     cache.WithConfig(config),
		//     cache.WithDatabase(driver.RedisDialectParam, redisConn),
		// )
		repo.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			hcheck.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)
		repo.attendance = attendance.New(
			attendance.WithConfig(config),
			attendance.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			attendance.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)
	}
}
