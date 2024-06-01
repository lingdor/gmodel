package mylog

import "log"

type Logger interface {
	Printf(format string, v ...any)
	Println(v ...any)
}

var defaultLogger Logger = log.Default()

func GetLogger() Logger {
	return defaultLogger
}

func SetLogger(logger Logger) {
	defaultLogger = logger
}
