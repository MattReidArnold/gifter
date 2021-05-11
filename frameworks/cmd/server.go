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
	"github.com/mattreidarnold/gifter/domain"
	log "github.com/mattreidarnold/gifter/frameworks/log"
	persistence "github.com/mattreidarnold/gifter/frameworks/persistence/inmem"
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

	httpAddr := ":8080"

	var kitLogger kitlog.Logger
	{
		kitLogger = kitlog.NewLogfmtLogger(os.Stderr)
		kitLogger = kitlog.With(kitLogger, "ts", kitlog.DefaultTimestampUTC)
		kitLogger = kitlog.With(kitLogger, "caller", kitlog.DefaultCaller)
	}

	logger := log.NewLogger(kitLogger)

	groupRepo := persistence.NewGroupRepository(logger, domain.NewGroup("1234", "group-name", 100))

	d := &app.Dependencies{
		Logger:          logger,
		GroupRepository: groupRepo,
	}

	var msgBus app.MessageBus
	{
		msgBus = app.NewMessageBus(logger)
		msgBus.RegisterCommandHandler(handlers.MakeAddGifter(d))
	}

	h := transport.MakeHTTPHandler(kitLogger, msgBus)

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
