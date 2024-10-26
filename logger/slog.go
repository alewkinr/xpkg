package logger

import (
	"fmt"
	"log/slog"

	"go.mrchanchal.com/zaphandler"
	"go.uber.org/zap"
)

// New — build a new slog.Logger with zap under the hood and sets this instant as a default one
func New(level string) *slog.Logger {
	zapL, _ := zap.NewProduction()
	defer func() { _ = zapL.Sync() }()

	logger := slog.New(zaphandler.New(zapL))
	slog.SetDefault(logger)

	var logLevel slog.Level
	if parseLogLevelErr := logLevel.UnmarshalText([]byte(level)); parseLogLevelErr != nil {
		panic(fmt.Errorf("failed to parse log level: %w", parseLogLevelErr))
	}

	slog.SetLogLoggerLevel(logLevel)

	return logger
}
