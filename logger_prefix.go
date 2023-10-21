package logger

import (
	"github.com/metadiv-io/logger/internal/constant"
	"github.com/metadiv-io/logger/internal/util"
)

func Prefix(prefix string) *prefixLogger {
	return &prefixLogger{prefix: prefix}
}

type prefixLogger struct {
	prefix string
}

// Info logs the information should be logged
func (l *prefixLogger) Info(v ...any) {
	util.LogToStd(LEVEL_INFO, l.prefix, v...)
	util.LogToFile(constant.LOG_FOLDER_PATH, util.GetLogFileName(), LEVEL_INFO, l.prefix, v...)
}

// Error logs the expected error
func (l *prefixLogger) Error(v ...any) {
	util.LogToStd(LEVEL_ERROR, l.prefix, v...)
	util.LogToFile(constant.LOG_FOLDER_PATH, util.GetLogFileName(), LEVEL_ERROR, l.prefix, v...)
}

// Debug logs the debug information, should be disabled in production
func (l *prefixLogger) Debug(v ...any) {
	if !constant.IS_DEBUG {
		return
	}
	util.LogToStd(LEVEL_DEBUG, l.prefix, v...)
}
