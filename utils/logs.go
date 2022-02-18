package utils

import (
	"log"
)

var LOG_LEVEL_DEBUG = 0
var LOG_LEVEL_WARNNING = 1
var LOG_LEVEL_ERROR = 2
var LOG_LEVEL_RELEASE = LOG_LEVEL_ERROR

// this is default log level.
var log_level = LOG_LEVEL_WARNNING

// There are 3 levels: DEBUG, WARNNING, ERROR
func SetLogLevel(level int) {
	log_level = level
}

func LogD(format string, v ...interface{}) {
	if log_level <= LOG_LEVEL_DEBUG {
		log.Printf("DEBUG: "+format, v...)
	}
}

func LogW(format string, v ...interface{}) {
	if log_level <= LOG_LEVEL_WARNNING {
		log.Printf("WARNING: "+format, v...)
	}
}

func LogE(format string, v ...interface{}) {
	if log_level <= LOG_LEVEL_ERROR {
		log.Printf("ERROR: "+format, v...)
	}
}

func LogFatal(format string, v ...interface{}) {
	log.Fatal("FATAL: "+format, v)
}
