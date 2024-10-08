package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang-service.codymj.io/config"
	"golang-service.codymj.io/db"
)

// Application is a struct to maintain application dependencies.
type application struct {
	cfg *config.Config
	wg  sync.WaitGroup
}

// Start starts the application.
func (a *application) start() {
	// Connect to database.
	db, err := db.New(a.cfg)
	if err != nil {
		log.Error().Msgf("failed to connect to database: %v", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Msgf("error closing database connection: %v", err)
		}
		log.Info().Msg("database connections closed")
	}()
	log.Info().Msg("database connection successful")

	// Server options.
	server := &http.Server{
		Addr:         a.cfg.Server.Host + ":" + a.cfg.Server.Port,
		Handler:      a.routes(db),
		ReadTimeout:  a.cfg.Server.Timeout.Read,
		WriteTimeout: a.cfg.Server.Timeout.Write,
		IdleTimeout:  a.cfg.Server.Timeout.Idle,
	}

	// Server shutdown error channel.
	shutdownErrorChan := make(chan error)

	// Listen for interrupts in a separate goroutine.
	go func() {
		// Channel to listen for interrupt signals.
		signalChan := make(chan os.Signal, 1)

		// Handle interrupts.
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		// Read signal from channel.
		s := <-signalChan
		log.Info().Msgf("stopping application: %s", s.String())

		// Context for graceful shutdown.
		ctx, cancel := context.WithTimeout(
			context.Background(),
			a.cfg.Server.Timeout.Server,
		)
		defer cancel()

		// Shutdown server.
		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErrorChan <- err
		}

		// Wait for cleanup tasks to finish.
		log.Info().Msg("completing background tasks...")
		a.wg.Wait()
		shutdownErrorChan <- nil
	}()

	// Startup.
	log.Info().Msg(fmt.Sprintf("server starting on %s", server.Addr))
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Error().Msgf("error starting server: %v", err)
		return
	}

	// Check for shutdown errors.
	err = <-shutdownErrorChan
	if err != nil {
		log.Error().Msgf("error shutting down server: %v", err)
		return
	}
}
