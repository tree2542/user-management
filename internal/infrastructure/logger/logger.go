package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	zerolog.Logger
}

func New() *Logger {
	// ใช้ console log ตอน dev, json log ตอน production
	var logger zerolog.Logger

	if os.Getenv("APP_ENV") == "production" {
		logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.Logger.Info().Fields(fields).Msg(msg)
}

func (l *Logger) Error(err error, msg string, fields ...interface{}) {
	l.Logger.Error().Err(err).Fields(fields).Msg(msg)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.Logger.Warn().Fields(fields).Msg(msg)
}