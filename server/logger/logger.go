package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
)

var logger *logrus.Logger

// Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// WithField
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// WithFields
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logger.WithFields(fields)
}

// WithError
func WithError(err error) *logrus.Entry {
	return logger.WithError(err)
}

func init() {

	// TODO finish config
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  "/var/log/info.log",
		logrus.ErrorLevel: "/var/log/error.log",
	}))
}
