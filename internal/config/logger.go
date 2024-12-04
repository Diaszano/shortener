// Package config provides utility functions and configuration management for the application,
// including setting up and managing the logging system.
package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Variables for the logger and synchronization.
// logger is a singleton instance of *zap.Logger.
// loggerOnce ensures the logger is initialized only once.
var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

// GetLogger initializes and returns the singleton instance of *zap.Logger.
// If the logger is already initialized, it returns the existing instance.
// The logger writes logs in JSON format to both the standard error and files in the "logs" directory.
// This function uses zap.Config to configure the logger, including log levels and encoding details.
func GetLogger() *zap.Logger {
	loggerOnce.Do(func() {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create logs directory: %v", err))
		}

		config := zap.Config{
			Level:       zap.NewAtomicLevelAt(Env.Server.GetLogLevel()),
			Development: Env.Server.IsDevelopment(),
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     customTimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stderr", "logs/app.log"},
			ErrorOutputPaths: []string{"stderr", "logs/error.log"},
		}

		// Build the logger based on the configuration.
		logger, err = config.Build()
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize logger: %v", err))
		}

		logger.Debug("Logger initialized successfully")
	})
	return logger
}

// CloseLogger flushes any buffered log entries and closes the logger.
// This should be called during application shutdown to ensure all logs are written properly.
func CloseLogger() {
	if logger != nil {
		_ = logger.Sync() // Ensure all buffered logs are written to their destinations.
	}
}

// customTimeEncoder defines a custom time format for log entries.
// It encodes time in RFC3339 format.
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}
