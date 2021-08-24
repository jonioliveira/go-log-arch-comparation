package injection

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	PanicLevel = "panic"
	ErrorLevel = "error"
	WarnLevel  = "warn"
	InfoLevel  = "info"
	DebugLevel = "debug"
)

// New creates a new `Logger`.
func New(level string) *logrus.Logger {
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
