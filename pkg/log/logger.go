package log

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Level - logging level.
type Level int

const (
	// LevelDebug indicates the level of debug messages logging.
	LevelDebug Level = iota
	// LevelInfo indicates the level of logging information messages.
	LevelInfo
	// LevelWarn indicates the logging level of warning messages.
	LevelWarn
	// LevelError indicates the logging level of error messages.
	LevelError
	// LevelFatal indicates the level of critical messages logging.
	LevelFatal
)

var currentLogLevel Level

// Init initializes the logger. Sets the output format.
func Init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Ldate)
	log.SetOutput(os.Stdout)
}

// SetLogLevel allows you to set the log level.
func SetLogLevel(logLevel Level) {
	currentLogLevel = logLevel
	Debug("Logging level is set:", currentLogLevel)
}

// String returns the debug message string.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "\033[36m\b\t[DEBUG]\t\033[0m"
	case LevelInfo:
		return "\033[37m\b\t[INFO]\t\033[0m"
	case LevelWarn:
		return "\033[33m\b\t[WARNING]\t\033[0m"
	case LevelError:
		return "\033[31m\b\t[ERROR]\t\033[0m"
	case LevelFatal:
		return "\033[35m\b\t[FATAL]\t\033[0m"
	default:
		return strconv.Itoa(int(l))
	}
}

func setOutput(level Level) {
	switch level {
	case LevelDebug:
		log.SetOutput(os.Stdout)
	case LevelInfo:
		log.SetOutput(os.Stdout)
	case LevelWarn:
		log.SetOutput(os.Stderr)
	case LevelError:
		log.SetOutput(os.Stderr)
	case LevelFatal:
		log.SetOutput(os.Stderr)
	default:
		log.SetOutput(os.Stdout)
	}
}

// Debugf allows you to print a debug message to stdout using a custom format.
func Debugf(format string, v ...interface{}) {
	if currentLogLevel > LevelDebug {
		return
	}
	setOutput(LevelDebug)
	_ = log.Output(2, fmt.Sprintf(LevelDebug.String()+format, v...))
}

// Infof allows you to print an informational message to stdout using a custom format.
func Infof(format string, v ...interface{}) {
	if currentLogLevel > LevelInfo {
		return
	}
	setOutput(LevelInfo)
	_ = log.Output(2, fmt.Sprintf(LevelInfo.String()+format, v...))
}

// Errorf allows you to print an error message to stdout using a custom format.
func Errorf(format string, v ...interface{}) {
	if currentLogLevel > LevelError {
		return
	}
	setOutput(LevelError)
	_ = log.Output(2, fmt.Sprintf(LevelError.String()+format, v...))
}

// Warningf allows you to print a warning message to stdout using a custom format.
func Warningf(format string, v ...interface{}) {
	if currentLogLevel > LevelWarn {
		return
	}

	setOutput(LevelWarn)
	_ = log.Output(2, fmt.Sprintf(LevelWarn.String()+format, v...))
}

// Fatalf allows you to print a fatal error message to stdout using a custom format.
// After displaying the message, stop the application with error code 1.
func Fatalf(format string, v ...interface{}) {
	if currentLogLevel > LevelFatal {
		return
	}
	setOutput(LevelFatal)
	_ = log.Output(2, fmt.Sprintf(LevelFatal.String()+format, v...))
	os.Exit(1)
}

// Debug allows you to print a debug message to stdout.
func Debug(v ...interface{}) {
	if currentLogLevel > LevelDebug {
		return
	}
	setOutput(LevelDebug)
	_ = log.Output(2, LevelDebug.String()+fmt.Sprintln(v...))
}

// Info allows you to print an informational message to stdout.
func Info(v ...interface{}) {
	if currentLogLevel > LevelInfo {
		return
	}
	setOutput(LevelInfo)
	_ = log.Output(2, LevelInfo.String()+fmt.Sprintln(v...))
}

// Error allows you to print an error message to stdout.
func Error(v ...interface{}) {
	if currentLogLevel > LevelError {
		return
	}
	setOutput(LevelError)
	_ = log.Output(2, LevelError.String()+fmt.Sprintln(v...))
}

// Warning allows you to print a warning message to stdout.
func Warning(v ...interface{}) {
	if currentLogLevel > LevelWarn {
		return
	}
	setOutput(LevelWarn)
	_ = log.Output(2, LevelWarn.String()+fmt.Sprintln(v...))
}

// Fatal allows you to print a fatal error message to stdout.
// After displaying the message, stop the application with error code 1.
func Fatal(v ...interface{}) {
	if currentLogLevel > LevelFatal {
		return
	}
	setOutput(LevelFatal)
	_ = log.Output(2, LevelFatal.String()+fmt.Sprintln(v...))
	os.Exit(1)
}
