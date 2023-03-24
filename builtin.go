package logger

import (
	"fmt"
	"log"
)

var (
	System  = wrap(&systemLogger{})
	Discard = wrap(&discardLogger{})
)

var logLevelNames = []string{
	"TRACE",
	"DEBUG",
	" INFO",
	" WARN",
	"ERROR",
}

type systemLogger struct{}

func (l *systemLogger) Log(level int, msg string, fields ...any) {
	l.LogArgs(level, nil, msg, fields...)
}

func (l *systemLogger) LogArgs(level int, args map[string]any, msg string, fields ...any) {
	if level < 0 || level >= len(logLevelNames) {
		return
	}
	if len(fields) == 0 {
		log.Printf("[%s] %s", logLevelNames[level], msg)
	} else {
		log.Printf("[%s] %s", logLevelNames[level], fmt.Sprintf(msg, fields...))
	}
}

type discardLogger struct{}

func (l *discardLogger) Log(level int, msg string, fields ...any)                          {}
func (l *discardLogger) LogArgs(level int, args map[string]any, msg string, fields ...any) {}
