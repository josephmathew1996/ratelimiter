package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeLogger sets up and returns a zap.SugaredLogger instance based on the provided log level.
func InitializeLogger(level string) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	var err error

	// Set up custom encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Use colored output
		EncodeTime:     zapcore.ISO8601TimeEncoder,      // Human-readable time format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Use console encoder instead of JSON
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Set the log level
	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	// Create a core with console encoder and log level
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), logLevel)

	// Build the logger
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
