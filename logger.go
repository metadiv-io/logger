package logger

import (
	"github.com/metadiv-io/logger/internal/constant"
	"github.com/metadiv-io/logger/internal/util"
)

// Info logs the information should be logged
func Info(v ...any) {
	util.LogToStd(LEVEL_INFO, "", v...)
	util.LogToFile(constant.LOG_FOLDER_PATH, util.GetLogFileName(), LEVEL_INFO, "", v...)
}

// Error logs the expected error
func Error(v ...any) {
	util.LogToStd(LEVEL_ERROR, "", v...)
	util.LogToFile(constant.LOG_FOLDER_PATH, util.GetLogFileName(), LEVEL_ERROR, "", v...)
}

// Debug logs the debug information, should be disabled in production
func Debug(v ...any) {
	if !constant.IS_DEBUG {
		return
	}
	util.LogToStd(LEVEL_DEBUG, "", v...)
}
