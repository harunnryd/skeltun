package driver

import (
	"fmt"
	"skeltun/config"
	"time"

	// main package.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
func OpenMysql(cfg config.IConfig) (dbx *sqlx.DB, err error) {
	opt := mysqlConfigurer(cfg)
	params := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", opt.Username, opt.Password, opt.Host, opt.Port, opt.DBName, opt.AddParams)
	dbx, err = sqlx.Open("mysql", params)
	if err != nil {
		return
	}

	dbx.SetConnMaxLifetime(time.Duration(opt.ConnMaxLifetime) * time.Second)
	dbx.SetMaxIdleConns(opt.MaxIdleConns)
	dbx.SetMaxOpenConns(opt.MaxOpenConns)
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
