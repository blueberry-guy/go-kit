package logger

import (
	"context"
	"fmt"

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

func (zl *zeroLogger) With(key string, value interface{}) Logger {
	var subLogger zerolog.Logger
	switch v := value.(type) {
	case string:
		subLogger = zl.logger.With().Str(key, v).Logger()
	case bool, int, int8, int16, int32, int64, float32, float64:
		subLogger = zl.logger.With().Str(key, fmt.Sprintf("%v", v)).Logger()
	default:
		subLogger = *zl.logger
	}
	return &zeroLogger{
		logger: &subLogger,
	}
}

func (zl *zeroLogger) WithContext(ctx context.Context) context.Context {
	return zl.logger.WithContext(ctx)
}
