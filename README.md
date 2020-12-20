<img alt="skeltun" src="https://i.imgur.com/62JGG0R.png" width="220"/>

This is a skeleton build with GO for our development process.

## Basic usage

**CLI**

```bash
foo@bar:~$ make docker-up
foo@bar:~$ make help
foo@bar:~$ make migration-sql create_users_table postgres
foo@bar:~$ make migrate-up postgres
foo@bar:~$ make migrate-down postgres
foo@bar:~$ make docker-down
```

**Migration files**

```bash
1481574547_create_users_table.up.sql
1481574547_create_users_table.down.sql
```


**UseCase**

> options.go

```go
package usecase

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase/hcheck"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"

	"github.com/gomodule/redigo/redis"
)

// Option ...
type Option func(*UseCase)

// WithDependency ...
func WithDependency(config config.IConfig) Option {
	var iRepo = repo.New(repo.WithDependency(config))
	var iPkg = pkg.New(pkg.WithDependency(config))
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(config.GetString("database.redis.hosts"))
		},
	}
	var iJob = job.New(job.WithConfig(config), job.WithRedis(redisPool))

	return func(usecase *UseCase) {
		// Inject all your UseCase's in here.
		// Example :
		// usecase.user = user.New(
		//    user.WithConfig(config),
		//    user.WithRepo(repo),
		// )
		usecase.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithRepo(iRepo),
			hcheck.WithPkg(iPkg),
			hcheck.WithJob(iJob),
		)
	}
}

```
> usecase.go

```go
package usecase

import (
	"github.com/harunnryd/skeltun/internal/app/usecase/hcheck"
)

// IUseCase ...
type IUseCase interface {
	// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
	GetHcheck() hcheck.IHcheck
}

// UseCase ...
type UseCase struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IUseCase {
	usecase := new(UseCase)
	for _, opt := range opts {
		opt(usecase)
	}
	return usecase
}

// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
func (usecase *UseCase) GetHcheck() hcheck.IHcheck {
	return usecase.hcheck
}

```

* **Repository**

> options.go

```go
package repo

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/repo/hcheck"
)

// Option ...
type Option func(*Repo)

// WithDependency ...
func WithDependency(config config.IConfig) Option {
	dbase := db.New(db.WithConfig(config))
	mysqlConn, _ := dbase.Manager(db.MysqlDialectParam)
	pgsqlConn, _ := dbase.Manager(db.PgsqlDialectParam)
	// onesignal := onesignal.New(onesignal.WithNetClient(&http.Client{
	// 	Timeout: time.Second * 10,
	// 	Transport: &http.Transport{
	// 		Dial: (&net.Dialer{
	// 			Timeout: 5 * time.Second,
	// 		}).Dial,
	// 		TLSHandshakeTimeout: 5 * time.Second,
	// 	},
	// }), onesignal.WithConfig(config))

	return func(repo *Repo) {
		// Inject all your repo's in here.
		// Example :
		// repo.cache = cache.New(
		//     cache.WithConfig(config),
		//     cache.WithDatabase(driver.RedisDialectParam, redisConn),
		// )s
		repo.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			hcheck.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)
	}
}

```
> repo.go

```go
package repo

import (
	"github.com/harunnryd/skeltun/internal/app/repo/hcheck"
)

// IRepo ...
type IRepo interface {
	// GetHcheck it returns instance of Hcheck that implements methods.
	GetHcheck() hcheck.IHcheck
}

// Repo ...
type Repo struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IRepo {
	repo := new(Repo)
	for _, opt := range opts {
		opt(repo)
	}
	return repo
}

// GetHcheck it returns instance of Hcheck that implements methods.
func (repo *Repo) GetHcheck() hcheck.IHcheck {
	return repo.hcheck
}

```

* **Handler**

> options.go

```go
package handler

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/hcheck"
	"github.com/harunnryd/skeltun/internal/app/usecase"
)

// Option ...
type Option func(*Handler)

// WithHandler ...
func WithHandler(config config.IConfig) Option {
	iUsecase := usecase.New(usecase.WithDependency(config))
	return func(handler *Handler) {
		// Inject all your handler's in here.
		// Example :
		// handler.user = user.New(
		//     user.WithConfig(config),
		//     user.WithUseCase(iUsecase),
		// )
		handler.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithUseCase(iUsecase),
		)
	}
}

```
> handler.go

```go
package handler

import (
	"github.com/harunnryd/skeltun/internal/app/handler/hcheck"
)

// IHandler ...
type IHandler interface {
	// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
	GetHcheck() hcheck.IHcheck
}

// Handler ...
type Handler struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IHandler {
	handler := new(Handler)
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
func (handler *Handler) GetHcheck() hcheck.IHcheck {
	return handler.hcheck
}

```

## Auxiliary packages

|  Package | Description  |
| ------------ | ------------ |
| [go-chi](https://github.com/go-chi/chi)  | Router (lightweight, idiomatic and composable router for building Go HTTP services) |
| [golang-migrate](https://github.com/golang-migrate/migrate)  | Database migrations  |
| [viper](https://github.com/spf13/viper)  | Go configuration with fangs  |
| [cobra](https://github.com/spf13/cobra)  | Go CLI  |
| [gorm](https://github.com/go-gorm/gorm)  | The fantastic ORM library for Golang, aims to be developer friendly  |

## Contributors
1. [Harun Nur Rasyid](https://github.com/harunnryd)
2. [Annahl Prayitno (Logo Design)](https://dribbble.com/AnboyStudio)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Copyright (c) 2020-present [Harun Nur Rasyid](https://github.com/harunnryd)

Licensed under [MIT License](./LICENSE)