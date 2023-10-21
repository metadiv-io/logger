package util

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/metadiv-io/logger/internal/types"
)

// GetStdLogger returns a logger that writes to stdout.
func GetStdLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

// GetFileLogger returns a logger that writes to a file.
func GetFileLogger(logFolderPath, filename string) *log.Logger {
	err := os.MkdirAll(logFolderPath, 0755)
	if err != nil {
		log.Fatalf("error creating log folder: %v", err)
	}
	path := logFolderPath + "/" + filename
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return log.New(file, "", log.LstdFlags)
}

// GetLogFileName returns a log file name in the format YYYY-MM-DD.log.
func GetLogFileName() string {
	return time.Now().Format("2006-01-02") + ".log"
}

// Log logs the information should be logged.
func Log(logger types.ILogger, level string, prefix string, v ...any) {
	if prefix != "" {
		logger.SetPrefix(fmt.Sprintf("[%s] %s ", level, prefix))
	} else {
		logger.SetPrefix(fmt.Sprintf("[%s] ", level))
	}
	logger.Println(v...)
}

// LogToStd logs the information should be logged to stdout.
func LogToStd(level string, prefix string, v ...any) {
	Log(GetStdLogger(), level, prefix, v...)
}

// LogToFile logs the information should be logged to a file.
func LogToFile(logFolderPath, filename string, level string, prefix string, v ...any) {
	Log(GetFileLogger(logFolderPath, filename), level, prefix, v...)
}
