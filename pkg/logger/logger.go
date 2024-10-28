package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// General method for raw logs
func (l *Logger) logNonFormatted(logger *log.Logger, v ...interface{}) {
	logger.Println(v...)
}

// General method for formatted logs
func (l *Logger) logFormatted(logger *log.Logger, format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) { l.logNonFormatted(l.debug, v...) }
func (l *Logger) Info(v ...interface{})  { l.logNonFormatted(l.info, v...) }
func (l *Logger) Warn(v ...interface{})  { l.logNonFormatted(l.warning, v...) }
func (l *Logger) Error(v ...interface{}) { l.logNonFormatted(l.err, v...) }

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) { l.logFormatted(l.debug, format, v...) }
func (l *Logger) Infof(format string, v ...interface{})  { l.logFormatted(l.info, format, v...) }
func (l *Logger) Warnf(format string, v ...interface{})  { l.logFormatted(l.warning, format, v...) }
func (l *Logger) Errorf(format string, v ...interface{}) { l.logFormatted(l.err, format, v...) }
