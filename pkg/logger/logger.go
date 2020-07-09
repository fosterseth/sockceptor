package logger

import (
	"log"
	"os"
)

var logLevel int
var showTrace bool

const (
	errorLevel = iota + 1
	warningLevel
	infoLevel
	debugLevel
)

// SetLogLevel is a helper function for setting logLevel int
func SetLogLevel(level int) {
	logLevel = level
}

// SetShowTrace is a helper function for setting showTrace bool
func SetShowTrace(trace bool) {
	showTrace = trace
}

// GetLogLevelByName is a helper function for returning level associated with log
// level string
func GetLogLevelByName(logName string) int {
	return LogLevelMap[logName]
}

// GetLogLevel returns current log level
func GetLogLevel() int {
	return logLevel
}

// LogLevelMap maps strings to log level int
// allows for --LogLevel Debug at command line
var LogLevelMap = map[string]int{
	"Error":   errorLevel,
	"Warning": warningLevel,
	"Info":    infoLevel,
	"Debug":   debugLevel,
}

// Error reports unexpected behavior, likely to result in termination
func Error(format string, v ...interface{}) {
	if logLevel >= errorLevel {
		log.SetPrefix("ERROR ")
		log.Printf(format, v...)
	}
}

// Warning reports unexpected behavior, not necessarily resulting in termination
func Warning(format string, v ...interface{}) {
	if logLevel >= warningLevel {
		log.SetPrefix("WARNING ")
		log.Printf(format, v...)
	}
}

// Info provides general purpose statements useful to end user
func Info(format string, v ...interface{}) {
	if logLevel >= infoLevel {
		log.SetPrefix("INFO ")
		log.Printf(format, v...)
	}
}

// Debug contains extra information helpful to developers
func Debug(format string, v ...interface{}) {
	if logLevel >= debugLevel {
		log.SetPrefix("DEBUG ")
		log.Printf(format, v...)
	}
}

// Trace outputs detailed packet traversal
func Trace(format string, v ...interface{}) {
	if showTrace {
		log.SetPrefix("TRACE ")
		log.Printf(format, v...)
	}
}

func init() {
	logLevel = infoLevel
	showTrace = false
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)
}
