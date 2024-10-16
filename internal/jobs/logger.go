package jobs

import (
	"fmt"
	stdlog "log"
	"os"
	"strings"

	"locgame-mini-server/pkg/log"
)

// Level - logging level.
type Level int

type JobLogger struct {
	currentLogLevel log.Level
	log             *stdlog.Logger
	output          *strings.Builder
}

// NewJobLogger initializes the job logger. Sets the output format.
func NewJobLogger() *JobLogger {
	l := &JobLogger{
		currentLogLevel: 0,
		output:          &strings.Builder{},
	}
	l.log = stdlog.New(l.output, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile|stdlog.Ldate)
	return l
}

// Debugf allows you to print a debug message to output using a custom format.
func (l *JobLogger) Debugf(format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(log.LevelDebug.String()+format, v...))
}

// Infof allows you to print an informational message to output using a custom format.
func (l *JobLogger) Infof(format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(log.LevelInfo.String()+format, v...))
}

// Errorf allows you to print an error message to output using a custom format.
func (l *JobLogger) Errorf(format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(log.LevelError.String()+format, v...))
}

// Warningf allows you to print a warning message to output using a custom format.
func (l *JobLogger) Warningf(format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(log.LevelWarn.String()+format, v...))
}

// Fatalf allows you to print a fatal error message to output using a custom format.
// After displaying the message, stop the application with error code 1.
func (l *JobLogger) Fatalf(format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(log.LevelFatal.String()+format, v...))
	os.Exit(1)
}

// Debug allows you to print a debug message to output.
func (l *JobLogger) Debug(v ...interface{}) {
	_ = l.log.Output(2, log.LevelDebug.String()+fmt.Sprintln(v...))
}

// Info allows you to print an informational message to output.
func (l *JobLogger) Info(v ...interface{}) {
	_ = l.log.Output(2, log.LevelInfo.String()+fmt.Sprintln(v...))
}

// Error allows you to print an error message to output.
func (l *JobLogger) Error(v ...interface{}) {
	_ = l.log.Output(2, log.LevelError.String()+fmt.Sprintln(v...))
}

// Warning allows you to print a warning message to output.
func (l *JobLogger) Warning(v ...interface{}) {
	_ = l.log.Output(2, log.LevelWarn.String()+fmt.Sprintln(v...))
}

// Fatal allows you to print a fatal error message to output.
// After displaying the message, stop the application with error code 1.
func (l *JobLogger) Fatal(v ...interface{}) {
	_ = l.log.Output(2, log.LevelFatal.String()+fmt.Sprintln(v...))
	os.Exit(1)
}
