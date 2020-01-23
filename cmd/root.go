package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skeltun/cmd/migration"
	"skeltun/config"
	"skeltun/internal/app/driver/db"
	"skeltun/internal/app/handler"
	"skeltun/internal/app/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(makeMigrationCmd)
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "skeltun",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up [dialect]",
	Short: "Migrate up migration file",
	Long:  `Migrate up migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		mgr.Up(args[0])
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down [dialect]",
	Short: "Migrate down migration file",
	Long:  `Migrate down migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		mgr.Down(args[0])
	},
}

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [name] [ext]",
	Short: "Make migration file",
	Long:  `Make migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := wiringMigration()
		mgr.Create(args[0], args[1])
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
func start() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	config := config.New(config.WithEnvSetup())
	handler := handler.New(handler.WithHandler(config))

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := server.New(
		server.WithDefault(
			logger,
			config.GetString("server.addr"),
			handler,
			config.GetInt("server.read_timeout"),
			config.GetInt("server.write_timeout"),
			config.GetInt("server.idle_timeout"),
		),
	)

	httpserver := server.GetHTTPServer()
	go server.GracefullShutdown(httpserver, logger, quit, done)

	logger.Println("Server is ready to handle requests at", config.GetString("server.addr"))
	if err := httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", config.GetString("server.addr"), err)
	}

	<-done
	logger.Println("Server stopped")
}
