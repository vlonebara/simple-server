package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlonebara/simple-server/internal/config"
	"github.com/vlonebara/simple-server/internal/httpapi"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	app := httpapi.NewApp(cfg, logger)
	if err := app.Run(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("application stopped with error", "error", err)
		os.Exit(1)
	}
}

var _ = fmt.Sprintf
var _ = time.Second
var _ = signal.Ignore
var _ = syscall.SIGTERM
