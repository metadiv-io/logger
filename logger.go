package logger

import "github.com/metadiv-io/env"

var LOG_FOLDER_PATH = env.String("LOG_FOLDER_PATH", "./logs")

const (
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_ERROR = "ERROR"
	LOG_LEVEL_FATAL = "FATAL"
	LOG_LEVEL_DEBUG = "DEBUG"
)

// Info logs the information should be logged
func Info(v ...any) {
	stdLogger := getStdLogger()
	fileLogger := getFileLogger(LOG_FOLDER_PATH, getLogFileName())
	print(stdLogger, LOG_LEVEL_INFO, "", v...)
	print(fileLogger, LOG_LEVEL_INFO, "", v...)
}

// Error logs the expected error
func Error(v ...any) {
	stdLogger := getStdLogger()
	fileLogger := getFileLogger(LOG_FOLDER_PATH, getLogFileName())
	print(stdLogger, LOG_LEVEL_ERROR, "", v...)
	print(fileLogger, LOG_LEVEL_ERROR, "", v...)
}

// Debug logs the debug information, should be disabled in production
func Debug(v ...any) {
	if env.String("GIN_MODE", "") == "release" {
		return
	}
	stdLogger := getStdLogger()
	print(stdLogger, LOG_LEVEL_DEBUG, "", v...)
}
