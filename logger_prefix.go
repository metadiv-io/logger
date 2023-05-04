package logger

import "github.com/metadiv-io/env"

func Prefix(prefix string) *prefixLogger {
	return &prefixLogger{prefix: prefix}
}

type prefixLogger struct {
	prefix string
}

// Info logs the information should be logged
func (l *prefixLogger) Info(v ...any) {
	stdLogger := getStdLogger()
	fileLogger := getFileLogger(LOG_FOLDER_PATH, getLogFileName())
	print(stdLogger, LOG_LEVEL_INFO, l.prefix, v...)
	print(fileLogger, LOG_LEVEL_INFO, l.prefix, v...)
}

// Error logs the expected error
func (l *prefixLogger) Error(v ...any) {
	stdLogger := getStdLogger()
	fileLogger := getFileLogger(LOG_FOLDER_PATH, getLogFileName())
	print(stdLogger, LOG_LEVEL_ERROR, l.prefix, v...)
	print(fileLogger, LOG_LEVEL_ERROR, l.prefix, v...)
}

// Debug logs the debug information, should be disabled in production
func (l *prefixLogger) Debug(v ...any) {
	if env.String("GIN_MODE", "") == "release" {
		return
	}
	stdLogger := getStdLogger()
	print(stdLogger, LOG_LEVEL_DEBUG, l.prefix, v...)
}
