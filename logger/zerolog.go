package logger

import (
	"github.com/rs/zerolog"
)

type zeroLogger struct {
	logger *zerolog.Logger
}

func (zl *zeroLogger) Info(format string, v ...interface{}) {
	zl.logger.Info().Msgf(format, v...)
}

func (zl *zeroLogger) Debug(format string, v ...interface{}) {
	zl.logger.Debug().Msgf(format, v...)
}

func (zl *zeroLogger) Warn(format string, v ...interface{}) {
	zl.logger.Warn().Msgf(format, v...)
}

func (zl *zeroLogger) Error(format string, v ...interface{}) {
	zl.logger.Error().Msgf(format, v...)
}

func (zl *zeroLogger) Print(format string, v ...interface{}) {
	zl.logger.Printf(format, v...)
}
