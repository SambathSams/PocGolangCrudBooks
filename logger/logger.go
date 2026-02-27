package logger

import (
	"fmt"
	apptime "go-crud-backend/util"
	"log"
	"path/filepath"
	"runtime"
)

func init() {
	// ignore default UTC time for printing logs - since we manually add app time
	log.SetFlags(0)
}

func Debug(message string, args ...interface{}) {
	logWithMetadata("DEBUG", message, args...)
}

func Info(message string, args ...interface{}) {
	logWithMetadata("INFO", message, args...)
}

func Imp(message string, args ...interface{}) {
	logWithMetadata("IMP", message, args...)
}

func Warn(message string, args ...interface{}) {
	logWithMetadata("WARN", message, args...)
}

// Error logs a message as an error
func Error(message string, args ...interface{}) {
	logWithMetadata("ERROR", message, args...)
}

func logWithMetadata(level string, message string, args ...interface{}) {
	// skip=2 goes back to the function that called Info() or Error()
	now := apptime.CurrentFormattedTime()
	pc, file, line, ok := runtime.Caller(2)

	details := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		// Get only the filename (not full path) and function name
		details = fmt.Sprintf("[%s:%d] %s()", filepath.Base(file), line, fn.Name())
	}

	formattedMessage := fmt.Sprintf(message, args...)

	// Final output: 2024/05/20 10:00:00 ERROR [main.go:42] main.main() -> Something went wrong
	log.Printf("%s %s %s -> %s", now, level, details, formattedMessage)
}
