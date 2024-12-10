package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs"
	sdk "github.com/agoda-com/opentelemetry-logs-go/sdk/logs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func newResource() *resource.Resource {
	hostName, err := os.Hostname()
	if err != nil {
		log.Printf("[WARN] Unable to retrieve hostname: %v", err)
		hostName = "unknown-host"
	}

	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("shortener"),
		semconv.ServiceVersionKey.String(Env.Server.Version),
		semconv.HostNameKey.String(hostName),
	)
}

func InitTracer() func() {
	traceProvider := trace.NewTracerProvider(
		trace.WithResource(newResource()),
	)

	otel.SetTracerProvider(traceProvider)

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := traceProvider.Shutdown(ctx); err != nil {
			log.Fatalf("[ERROR] TracerProvider shutdown failed: %v", err)
		}
		log.Println("[INFO] Tracer shutdown completed")
	}
}

func InitLogger() (*sdk.LoggerProvider, func()) {
	ctx := context.Background()

	logExporter, err := otlplogs.NewExporter(ctx)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize log exporter: %v", err)
	}

	loggerProvider := sdk.NewLoggerProvider(
		sdk.WithBatcher(logExporter),
		sdk.WithResource(newResource()),
	)

	return loggerProvider, func() {
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := logExporter.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("[ERROR] Failed to shut down log exporter: %v", err)
		}
		log.Println("[INFO] Logger shutdown completed")
	}
}

// SetupLogger configures and returns a new zap.Logger instance that integrates
// OpenTelemetry logging with a JSON encoder and standard output. The logger uses
// the logging level specified in the application configuration.
func SetupLogger(loggerProvider *sdk.LoggerProvider) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		otelzap.NewOtelCore(loggerProvider),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(Env.Server.GetLogLevel()),
		),
	)
	return zap.New(core)
}
