package main

import (
	"github.com/pablogolobaro/chequery/internal/app"
	"github.com/pablogolobaro/chequery/internal/config"
	_ "github.com/pablogolobaro/chequery/internal/docs"
	"github.com/pablogolobaro/chequery/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logger.Get()

	logger.Debug("Loading config...")
	conf := config.Load()

	logger.Debug("Loaded config...")

	application := app.NewApplication(logger)

	logger.Info("Bootstrap Application")

	err := application.Bootstrap(conf)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("RegisterRouter Application")

	err = application.RegisterRouter()
	if err != nil {
		logger.Fatal(err)
	}

	chStop := make(chan os.Signal)
	errCh := make(chan error)

	signal.Notify(chStop, syscall.SIGTERM)

	go func() {
		logger.Info("Start Application")
		err := application.Start(conf)
		errCh <- err
	}()

	select {
	case <-chStop:
		logger.Info("Stop Application")

		err = application.Stop()
		if err != nil {
			logger.Errorw("Cannot stop Application gracefully", "error:", err)
		}
	case err = <-errCh:
		logger.Errorw("Application error", "error:", err)
	}

	logger.Info("Application Stopped")
}
