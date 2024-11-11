package logger

import (
	"go.uber.org/zap"
)

var (
	rawLogger *zap.Logger        // Regular zap.Logger instance
	sugar     *zap.SugaredLogger // Sugared logger instance
)

// init initializes both zap logger instances as global variables
func init() {
	var err error
	rawLogger, err = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize zap logger: " + err.Error())
	}
	sugar = rawLogger.Sugar()
}

// GetLogger returns the global zap.Logger instance (non-sugared)
func GetLogger() *zap.Logger {
	return rawLogger
}

// GetSugaredLogger returns the global zap.SugaredLogger instance
func GetSugaredLogger() *zap.SugaredLogger {
	return sugar
}

// Sync flushes any buffered log entries. Call this at the end of main() to ensure logs are written.
func Sync() {
	if rawLogger != nil {
		_ = rawLogger.Sync()
	}
}
