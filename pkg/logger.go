package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DEBUG = "\033[1;36mDEBUG\033[0m"
	INFO  = "\033[1;34mINFO\033[0m"
	WARN  = "\033[1;33mWARN\033[0m"
	ERROR = "\033[1;31mERROR\033[0m"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

type Logger struct {
	Module         string
	ActiveLogLevel LogLevel
	logger         *log.Logger
}

func NewLogger(module string, activeLogLevel LogLevel) *Logger {
	return &Logger{
		Module:         module,
		ActiveLogLevel: activeLogLevel,
		logger:         log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) logMessage(level string, levelEnum LogLevel, format string, args ...interface{}) {
	if levelEnum < l.ActiveLogLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	logLine := fmt.Sprintf("[%s] [%s] [%s]: %s", timestamp, level, l.Module, message)

	l.logger.Println(logLine)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.logMessage(DEBUG, Debug, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.logMessage(INFO, Info, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.logMessage(WARN, Warn, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.logMessage(ERROR, Error, format, args...)
}
