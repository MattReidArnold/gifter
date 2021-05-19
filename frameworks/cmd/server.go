package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/app/handlers"
	"github.com/mattreidarnold/gifter/frameworks/config"
	"github.com/mattreidarnold/gifter/frameworks/id"
	log "github.com/mattreidarnold/gifter/frameworks/log"
	"github.com/mattreidarnold/gifter/frameworks/persistence/mongo"
	"github.com/mattreidarnold/gifter/frameworks/transport"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Run:   serverRun,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverRun(cmd *cobra.Command, args []string) {

	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	httpAddr := ":8080"

	var kitLogger kitlog.Logger
	{
		kitLogger = kitlog.NewLogfmtLogger(os.Stderr)
		kitLogger = kitlog.With(kitLogger, "ts", kitlog.DefaultTimestampUTC)
		kitLogger = kitlog.With(kitLogger, "caller", kitlog.DefaultCaller)
	}

	logger := log.NewLogger(kitLogger)

	mongoConn := mongo.Connection{
		Database: config.MongoDatabase,
		Host:     config.MongoHost,
		Password: config.MongoPassword,
		Port:     config.MongoPort,
		Username: config.MongoUsername,
	}
	mongoClient, disconnect, err := mongo.NewClient(logger, mongoConn)
	if err != nil {
		kitLogger.Log("msg", "Error creating mongo client", "err", err)
		return
	}
	defer disconnect()

	groupRepo := mongo.NewGroupRepository(mongoClient)

	msgBus := app.NewMessageBus(logger)

	d := &app.Dependencies{
		Logger:          logger,
		GroupRepository: groupRepo,
		MessageBus:      msgBus,
		GenerateID:      id.GenerateId,
	}

	handlers.RegisterAll(d)

	h := transport.MakeHTTPHandler(kitLogger, d)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		kitLogger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	kitLogger.Log("exit", <-errs)
}
