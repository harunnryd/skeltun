// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/harunnryd/skeltun/cmd/listener"
	"github.com/harunnryd/skeltun/cmd/migration"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/handler"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/server"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

const (
	asciiArt = `
    __          _ __         
   / /_  ____  (_) /__  _____
  / __ \/ __ \/ / / _ \/ ___/
 / /_/ / /_/ / / /  __/ /    
/_.___/\____/_/_/\___/_/     PLATE
`
)

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(makeMigrationCmd)
	rootCmd.AddCommand(routeListCmd)
	rootCmd.AddCommand(workerCmd)
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "skeltun",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		doStart()
	},
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start Worker",
	Long:  "Start Worker",
	Run: func(cmd *cobra.Command, args []string) {
		doWorker()
	},
}

var routeListCmd = &cobra.Command{
	Use:   "route:list",
	Short: "Route list",
	Long:  "Route list",
	Run: func(cmd *cobra.Command, args []string) {
		getRoutes()
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up [dialect]",
	Short: "Migrate up migration file",
	Long:  `Migrate up migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		if err := mgr.Up(args[0]); err != nil {
			fmt.Printf("Migrate up error: %v\n", err.Error())
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down [dialect]",
	Short: "Migrate down migration file",
	Long:  `Migrate down migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		if err := mgr.Down(args[0]); err != nil {
			fmt.Printf("Migrate down error: %s", err.Error())
		}
	},
}

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [name] [ext]",
	Short: "Make migration file",
	Long:  `Make migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		_ = mgr.Create(args[0], args[1])
	},
}

// Execute executes the root command.
func Execute() (err error) {
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func wiringMigration() (mgr migration.IMigration) {
	cfg := config.New(config.WithEnvSetup())
	dbase := db.New(db.WithConfig(cfg))
	mysqlConn, _ := dbase.Manager(db.MysqlDialectParam)
	pgsqlConn, _ := dbase.Manager(db.PgsqlDialectParam)

	mgr = migration.New(
		// Wiring your option in here
		migration.WithDatabase(db.MysqlDialectParam, mysqlConn),
		migration.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
	)
	return
}

// based on article https://marcofranssen.nl/go-webserver-with-graceful-shutdown/
func doStart() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	iConfig := config.New(config.WithEnvSetup())
	iHandler := handler.New(handler.WithHandler(iConfig))

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	iServer := server.New(
		server.WithDefault(
			logger,
			iConfig.GetString("server.addr"),
			iHandler,
			iConfig.GetInt("server.read_timeout"),
			iConfig.GetInt("server.write_timeout"),
			iConfig.GetInt("server.idle_timeout"),
		),
	)

	httpserver := iServer.GetHTTPServer()
	go iServer.GracefullShutdown(httpserver, logger, quit, done)

	logger.Print(asciiArt)
	logger.Println("=> http server started on", iConfig.GetString("server.addr"))
	if err := httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", iConfig.GetString("server.addr"), err)
	}

	<-done
	logger.Println("Server stopped")
}

func doWorker() {
	var cfg = config.New(config.WithEnvSetup())
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(cfg.GetString("database.redis.hosts"))
		},
	}
	var iRepo = repo.New(repo.WithDependency(cfg))
	var iUseCase = usecase.New(usecase.WithDependency(cfg))
	var iPkg = pkg.New(pkg.WithDependency(cfg))
	var iJob = job.New(job.WithConfig(cfg), job.WithRedis(redisPool))

	listener.New(
		listener.WithConfig(cfg),
		listener.WithRedis(redisPool),
		listener.WithRepo(iRepo),
		listener.WithUseCase(iUseCase),
		listener.WithPkg(iPkg),
		listener.WithJob(iJob),
	).Start()
}

func getRoutes() {
	logger := log.New(os.Stdout, "route: ", log.LstdFlags)
	cfg := config.New(config.WithEnvSetup())
	h := handler.New(handler.WithHandler(cfg))
	w := server.New().Router(h)

	logger.Print(asciiArt)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Method", "Route"})

	// debug
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		table.Append([]string{method, route})
		return nil
	}

	if err := chi.Walk(w, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	table.Render() // Send output
}
