package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	baseLogger *logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger(testEnv bool) *StandardLogger {

	var baseLogger = logrus.New()
	baseLogger.Level = logrus.ErrorLevel
	if !testEnv {

		baseLogger.Out = ioutil.Discard
	}
	baseLogger.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		ForceQuote:      false,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	return &StandardLogger{baseLogger}
}

// Info is a standard info logger
func (sl *StandardLogger) Info(format string, v ...interface{}) {

	sl.baseLogger.Infof(format, v...)
}

// Warn is a standard warning logger
func (sl *StandardLogger) Warn(format string, v ...interface{}) {

	sl.baseLogger.Warnf(format, v...)
}

// Error is a standard error logger
func (sl *StandardLogger) Error(format string, v ...interface{}) {

	sl.baseLogger.Errorf(format, v...)
}

// Debug is a standard debug logger
func (sl *StandardLogger) Debug(format string, v ...interface{}) {

	sl.baseLogger.Debugf(format, v...)
}

// SetOutput permit change the Logger Out
func (sl *StandardLogger) SetOutput(out io.Writer) {

	if out != nil {
		sl.baseLogger.Out = out
	}
}

// SetLevel permit change the Logger Level
func (sl *StandardLogger) SetLevel(level logrus.Level) {

	if level != sl.baseLogger.Level {
		sl.baseLogger.Level = level
	}
}
