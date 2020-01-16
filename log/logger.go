package log

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fsufitch/wire-web-demo/config"
)

// Level is an enum of log levels
type Level int

// Values for LogLevel
const (
	Debug Level = 0 + iota
	Info
	Warning
	Error
	Critical
)

// MultiLogger provides a three tiered (ignore, print, error) logger wrapper around log.MultiLogger
type MultiLogger struct {
	PrintLogger *log.Logger
	ErrorLogger *log.Logger
	PrintLevel  Level
	ErrorLevel  Level
}

// SetPrefix sets the prefix for both wrapped loggers
func (l MultiLogger) SetPrefix(prefix string) {
	l.PrintLogger.SetPrefix(prefix)
	l.ErrorLogger.SetPrefix(prefix)
}

// SetFlags sets the flags for both wrapped loggers
func (l MultiLogger) SetFlags(flags int) {
	l.PrintLogger.SetFlags(flags)
	l.ErrorLogger.SetFlags(flags)
}

// SetOutput sets the outputs for the print and error output loggers
func (l MultiLogger) SetOutput(printOutput io.Writer, errorOutput io.Writer) {
	l.PrintLogger.SetOutput(printOutput)
	l.ErrorLogger.SetOutput(errorOutput)
}

func (l MultiLogger) printf(level Level, format string, values ...interface{}) error {
	var activeLogger *log.Logger
	if level < l.PrintLevel {
		return nil
	} else if level < l.ErrorLevel {
		activeLogger = l.PrintLogger
	} else {
		activeLogger = l.ErrorLogger
	}

	var levelPrefix string
	switch level {
	case Debug:
		levelPrefix = "[DEBUG]"
	case Info:
		levelPrefix = "[INFO]"
	case Warning:
		levelPrefix = "[WARNING]"
	case Error:
		levelPrefix = "[ERROR]"
	case Critical:
		levelPrefix = "[CRITICAL]"
	}

	message := fmt.Sprintf("%s %s", levelPrefix, fmt.Sprintf(format, values...))
	return activeLogger.Output(3, message)
}

// Debugf prints a debug message to the appropriate logger
func (l MultiLogger) Debugf(format string, values ...interface{}) {
	l.printf(Debug, format, values...)
}

// Infof prints a debug message to the appropriate logger
func (l MultiLogger) Infof(format string, values ...interface{}) {
	l.printf(Info, format, values...)
}

// Warningf prints a debug message to the appropriate logger
func (l MultiLogger) Warningf(format string, values ...interface{}) {
	l.printf(Warning, format, values...)
}

// Errorf prints a debug message to the appropriate logger
func (l MultiLogger) Errorf(format string, values ...interface{}) {
	l.printf(Error, format, values...)
}

// Criticalf prints a debug message to the appropriate logger
func (l MultiLogger) Criticalf(format string, values ...interface{}) {
	l.printf(Critical, format, values...)
}

func (l MultiLogger) flush() {
	l.PrintLogger.Writer().Write([]byte("\n"))
	l.ErrorLogger.Writer().Write([]byte("\n"))
}

// ProvideStdOutErrMultiLogger creates a custom MultiLogger based on whether we are running in debug mode
func ProvideStdOutErrMultiLogger(debugMode config.DebugMode) (*MultiLogger, func()) {
	stdOutLevel := Info
	stdErrLevel := Error
	flags := log.Ldate | log.Ltime | log.LUTC

	if debugMode {
		stdOutLevel = Debug
		flags |= log.Llongfile | log.Lmicroseconds
	}

	logger := &MultiLogger{
		PrintLevel:  stdOutLevel,
		ErrorLevel:  stdErrLevel,
		PrintLogger: log.New(os.Stdout, "", flags),
		ErrorLogger: log.New(os.Stdout, "", flags),
	}

	logger.Infof("initialized logging")

	return logger, func() { logger.flush() }
}
