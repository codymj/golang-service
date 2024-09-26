package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang-service.codymj.io/config"
	"golang-service.codymj.io/db"
)

// Application is a struct to maintain application dependencies.
type application struct {
	cfg    *config.Config
	logger *slog.Logger
	wg     *sync.WaitGroup
}

// Run starts the HTTP server.
func (a *application) serve() {
	// Server options.
	server := &http.Server{
		Addr:         a.cfg.Server.Host + ":" + a.cfg.Server.Port,
		Handler:      nil,
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

		s := <-signalChan
		a.logger.Info(fmt.Sprintf("caught signal: %s", s.String()))

		// Context for graceful shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.Timeout.Server)
		defer cancel()

		// Handle interrupts.
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		// Shutdown server.
		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErrorChan <- err
		}

		a.logger.Info("completing background tasks")
		a.wg.Wait()
		shutdownErrorChan <- nil
	}()

	// Connect to database.
	db, err := db.New(a.cfg)
	if err != nil {
		a.logger.Error(fmt.Sprintf("failed to connect to database: %v", err))
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			a.logger.Error(fmt.Sprintf("error closing database connection: %v", err))
		}
	}()
	a.logger.Info("database connection successful")

	// Startup.
	a.logger.Info(fmt.Sprintf("server starting on %s", server.Addr))
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		a.logger.Error(fmt.Sprintf("error starting server: %v", err))
		return
	}

	err = <-shutdownErrorChan
	if err != nil {
		a.logger.Error(fmt.Sprintf("error shutting down server: %v", err))
		return
	}

	a.logger.Info(fmt.Sprintf("stopped server: %s", server.Addr))
}
