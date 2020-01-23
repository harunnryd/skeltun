package driver

import (
	"fmt"
	"skeltun/config"

	// main package.
	"github.com/jmoiron/sqlx"
	// main package.
	_ "github.com/lib/pq"
)

// PgsqlOption ...
type PgsqlOption struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	MaxPoolSize int
	Sslmode     string
}

// OpenPgsql ...
func OpenPgsql(cfg config.IConfig) (dbx *sqlx.DB, err error) {
	opt := pgsqlConfigurer(cfg)
	params := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", opt.Host, opt.Port, opt.Username, opt.DBName, opt.Password, opt.Sslmode)
	dbx, err = sqlx.Open("postgres", params)
	if err != nil {
		return
	}

	dbx.SetMaxOpenConns(opt.MaxPoolSize)
	return
}

func pgsqlConfigurer(cfg config.IConfig) PgsqlOption {
	return PgsqlOption{
		Host:        cfg.GetString("database.pgsql.host"),
		Port:        cfg.GetInt("database.pgsql.port"),
		Username:    cfg.GetString("database.pgsql.username"),
		Password:    cfg.GetString("database.pgsql.password"),
		DBName:      cfg.GetString("database.pgsql.db_name"),
		MaxPoolSize: cfg.GetInt("database.pgsql.max_pool_size"),
		Sslmode:     cfg.GetString("database.pgsql.sslmode"),
	}
}
