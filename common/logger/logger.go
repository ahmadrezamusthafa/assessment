package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	LogLevelTrace   = "TRACE"
	LogLevelDebug   = "DEBUG"
	LogLevelInfo    = "INFO"
	LogLevelWarning = "WARNING"
	LogLevelError   = "ERROR"
	LogLevelPanic   = "PANIC"
)

var logLevel = map[string]int{
	LogLevelTrace:   5,
	LogLevelDebug:   4,
	LogLevelInfo:    3,
	LogLevelWarning: 2,
	LogLevelError:   1,
	LogLevelPanic:   0,
}

var ptSystemName string

var activeLogLevel = strings.ToUpper(os.Getenv("LOG_LEVEL"))

func parseLogLevel() string {
	switch activeLogLevel {
	case LogLevelTrace, LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError, LogLevelPanic:
	default:
		activeLogLevel = LogLevelTrace
	}
	return activeLogLevel
}

func getActiveLogLevel() int {
	return logLevel[activeLogLevel]
}

func SetupLogger() {
	activeLogLevel = parseLogLevel()
	log.SetPrefix("")
	log.SetFlags(0)
}

func Warn(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelWarning] {
		message := fmt.Sprintf("WARN: "+format, v...)
		log.Print(message)
	}
}

func Trace(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelTrace] {
		message := fmt.Sprintf("TRACE: "+format, v...)
		log.Print(message)
	}
}

func Debug(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelDebug] {
		message := fmt.Sprintf("DEBUG: "+format, v...)
		log.Print(message)
	}
}

func Info(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelInfo] {
		message := fmt.Sprintf("INFO: "+format, v...)
		log.Print(message)
	}
}

func Err(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelError] {
		message := []interface{}{fmt.Sprintf("ERROR: "+format, v...)}
		log.Print(message...)
	}
}

func Panic(format string, v ...interface{}) {
	if getActiveLogLevel() >= logLevel[LogLevelPanic] {
		message := []interface{}{fmt.Sprintf("PANIC: "+format, v...)}
		log.Print(message...)
	}
}

func Fatal(format string, v ...interface{}) {
	Err(format, v...)
	os.Exit(1)
}
