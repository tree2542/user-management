package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"user-management/internal/usecase/port"
)

var _ port.Logger = (*Logger)(nil)

type Logger struct {
	l zerolog.Logger
}

func New() *Logger {

	// dev mode = console log
	if os.Getenv("APP_ENV") != "production" {
		consoleWriter := zerolog.ConsoleWriter{
			Out: os.Stdout,
		}

		return &Logger{
			l: log.Output(consoleWriter),
		}
	}

	// production = JSON log
	return &Logger{
		l: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (lg *Logger) Info(msg string, fields ...interface{}) {
	lg.l.Info().Fields(fields).Msg(msg)
}

func (lg *Logger) Error(msg string, fields ...interface{}) {
	lg.l.Error().Fields(fields).Msg(msg)
}

func (lg *Logger) Warn(msg string, fields ...interface{}) {
	lg.l.Warn().Fields(fields).Msg(msg)
}
