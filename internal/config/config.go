package config

import (
	"log/slog"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPAddr        string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	LogLevel        slog.Level
}

func Load() Config {
	return Config{
		HTTPAddr:        env("HTTP_ADDR", ":8080"),
		ReadTimeout:     durationEnv("HTTP_READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    durationEnv("HTTP_WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     durationEnv("HTTP_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: durationEnv("SHUTDOWN_TIMEOUT", 10*time.Second),
		LogLevel:        logLevel(env("LOG_LEVEL", "info")),
	}
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func durationEnv(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func logLevel(value string) slog.Level {
	if level, err := strconv.Atoi(value); err == nil {
		return slog.Level(level)
	}
	switch value {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
