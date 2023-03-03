package main

import (
	"github.com/pablogolobaro/chequery/internal/app"
	"github.com/pablogolobaro/chequery/internal/config"
	"github.com/pablogolobaro/chequery/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.Get()

	log.Debug("Loading config...")
	conf := config.Load()

	log.Debug("Loaded config...")

	application := app.NewApplication(log)

	log.Info("Bootstrap Application")

	err := application.Bootstrap(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("RegisterRouter Application")

	err = application.RegisterRouter()
	if err != nil {
		log.Fatal(err)
	}

	chStop := make(chan os.Signal)
	errCh := make(chan error)

	signal.Notify(chStop, syscall.SIGTERM)

	go func() {
		log.Info("Start Application")
		err := application.Start(conf)
		errCh <- err
	}()

	select {
	case <-chStop:
		log.Info("Stop Application")

		err = application.Stop()
		if err != nil {
			log.Errorw("Cannot stop Application gracefully", "error:", err)
		}
	case err = <-errCh:
		log.Errorw("Application error", "error:", err)
	}

	log.Info("Application Stopped")
}
