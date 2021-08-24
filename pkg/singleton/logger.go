package singleton

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log    *logrus.Logger
	doOnce sync.Once
)

const (
	PanicLevel = "panic"
	ErrorLevel = "error"
	WarnLevel  = "warn"
	InfoLevel  = "info"
	DebugLevel = "debug"
)

func GetInstance() *logrus.Logger {
	doOnce.Do(func() {
		log = new(DebugLevel)
	})
	return log
}

// New creates a new `Logger`.
func new(level string) *logrus.Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	if l, err := logrus.ParseLevel(level); err != nil {
		log.WithField("invalidLevel", level).
			Error("invalid log level, defaulting to 'info'")
	} else {
		log.SetLevel(l)
		log.WithField("to", level).
			Info("log level set")
	}
	return log
}

// Debug logs a entry with the debug level.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logs a entry with the info level.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logs a entry with the warn level.
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logs a entry with the error level.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Panic logs a entry with the panic level and then calls panic.
// WARN: calling this will terminate the process.
func Panic(args ...interface{}) {
	log.Panic(args...)
}
