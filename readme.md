# skeltun

This is a skeleton build with GO for our development process.

---

## Basic usage

* **CLI**

```{r, engine='bash', count_lines, properties}
foo@bar:~$ go run main.go help
foo@bar:~$ go run main.go make:migration create_users_table postgres
foo@bar:~$ go run main.go migrate:up postgres
foo@bar:~$ go run main.go migrate:down postgres
```

* **Service**

`option.go`

```go
package service

import (
  "skeltun/config"
  "skeltun/internal/app/repository"
  "skeltun/internal/app/service/hcheck"
)

// Option ...
type Option func(*Service)

// WithService ...
func WithService(config config.IConfig) Option {
 repo := repository.New(repository.WithDatabase(config))

  return func(svc *Service) {
      // Inject all your service's in here.
      // Example :
      // svc.user = user.New(
      //    user.WithConfig(config),
      //    user.WithRepo(repo),
      // )
      svc.hcheck = hcheck.New(
      hcheck.WithConfig(config),
      hcheck.WithRepo(repo),
    )
  }
}
```

`service.go`

```go
package service

import (
  "skeltun/internal/app/service/hcheck"
)

// IService ...
type IService interface {
  Hcheck() hcheck.IHcheck
  // User() user.IUser
}

// Service ...
type Service struct {
  hcheck hcheck.IHcheck
  // user user.IUser
}

// New ...
func New(opts ...Option) IService {
  svc := new(Service)
  for _, opt := range opts {
    opt(svc)
  }
  return svc
}

// Hcheck ...
func (svc *Service) Hcheck() hcheck.IHcheck {
  return svc.hcheck
}

// User ...
// func (svc *Service) User() user.IUser {
//    return svc.user
// }
```

* **Repository**

`option.go`

```go
package repository

import (
  "skeltun/config"
  "skeltun/internal/app/driver/db"
  "skeltun/internal/app/repository/hcheck"
)

// Option ...
type Option func(*Repository)

// WithDatabase ...
func WithDatabase(config config.IConfig) Option {
  dbase        := db.New(db.WithConfig(config))
  mysqlConn, _ := dbase.Manager(db.MysqlDialectParam)
  pgsqlConn, _ := dbase.Manager(db.PgsqlDialectParam)

  return func(repo *Repository) {
    // Inject all your repo's in here.
    // Example :
    // repo.user = user.New(
    //   user.WithConfig(config),
    //   user.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
    //   user.WithDatabase(db.MysqlDialectParam, mysqlConn),
    // )
    repo.hcheck = hcheck.New(
      hcheck.WithConfig(config),
      hcheck.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
      hcheck.WithDatabase(db.MysqlDialectParam, mysqlConn),
    )
  }
}
```

`repo.go`

```go
package repository

import (
  "skeltun/internal/app/repository/hcheck"
)

// IRepository ...
type IRepository interface {
  Hcheck() hcheck.IHcheck
  // User() user.IUser
}

// Repository ...
type Repository struct {
  hcheck hcheck.IHcheck
  // user user.IUser
}

// New ...
func New(opts ...Option) IRepository {
  repo := new(Repository)
  for _, opt := range opts {
    opt(repo)
  }
  return repo
}

// Hcheck ...
func (repo *Repository) Hcheck() hcheck.IHcheck {
  return repo.hcheck
}

// User ...
// func (repo *Repository) User() user.IUser {
//   return repo.user
// }
```

* **Handler**

`option.go`

```go
package handler

import (
  "skeltun/config"
  "skeltun/internal/app/handler/hcheck"
  "skeltun/internal/app/service"
)

// Option ...
type Option func(*Handler)

// WithHandler ...
func WithHandler(config config.IConfig) Option {
  service := service.New(service.WithService(config))
  return func(handler *Handler) {
    // Inject all your handler's in here.
    // Example :
    // handler.user = user.New(
    //     user.WithConfig(config),
    //     user.WithService(service),
    // )
    handler.hcheck = hcheck.New(
      hcheck.WithConfig(config),
      hcheck.WithService(service),
    )
  }
}
```

`handler.go`

```go
package handler

import "skeltun/internal/app/handler/hcheck"

// IHandler ...
type IHandler interface {
  Hcheck() hcheck.IHcheck
  // User() user.IUser
}

// Handler ...
type Handler struct {
  hcheck hcheck.IHcheck
  // user user.IUser
}

// New ...
func New(opts ...Option) IHandler {
  handler := new(Handler)
  for _, opt := range opts {
    opt(handler)
  }
  return handler
}

// Hcheck ...
func (handler *Handler) Hcheck() hcheck.IHcheck {
  return handler.hcheck
}

// User ...
// func (handler *Handler) User() user.IUser {
//    return user.user
// }
```

---

## Migration files

Each migration has an up and down migration.

```{r, engine='bash', count_lines, properties}
1481574547_create_users_table.up.sql
1481574547_create_users_table.down.sql
```

---

## Auxiliary packages

| Package                                            | Description                                                 |
|:---------------------------------------------------|:-------------------------------------------------------------
| [go-chi](https://github.com/go-chi/chi)             | Router (lightweight, idiomatic and composable router for building Go HTTP services)                       |
| [golang-migrate](https://github.com/golang-migrate/migrate)         | Database migrations                         |
| [cobra](https://github.com/spf13/cobra)       | Go CLI                                          |
| [viper](https://github.com/spf13/viper) | Go configuration with fangs                        |
| [sqlx](https://github.com/jmoiron/sqlx)   | General purpose extensions to golang's database/sql

---

## License

Copyright (c) 2020-present [Harun Nur Rasyid](https://github.com/harunnryd)

Licensed under [MIT License](./LICENSE)
