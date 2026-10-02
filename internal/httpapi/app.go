package httpapi

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vlonebara/simple-server/internal/config"
)

type App struct {
	config config.Config
	logger *slog.Logger
}

func NewApp(cfg config.Config, logger *slog.Logger) *App {
	return &App{config: cfg, logger: logger}
}

func (a *App) Run(parent context.Context) error {
	server := &http.Server{
		Addr:              a.config.HTTPAddr,
		Handler:           a.routes(),
		ReadTimeout:       a.config.ReadTimeout,
		WriteTimeout:      a.config.WriteTimeout,
		IdleTimeout:       a.config.IdleTimeout,
		ReadHeaderTimeout: a.config.ReadTimeout,
	}

	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errs := make(chan error, 1)
	go func() {
		a.logger.Info("HTTP server started", "address", a.config.HTTPAddr)
		errs <- server.ListenAndServe()
	}()

	select {
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("HTTP server error: %w", err)
	case <-ctx.Done():
		a.logger.Info("graceful shutdown started")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}
	a.logger.Info("HTTP server stopped")
	return nil
}
