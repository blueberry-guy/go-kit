package logger

import (
	"os"

	"github.com/blueberry-guy/go-kit/env"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New returns logger based on env. type
// this is an opinionated approach for choosing the logger configurations
// based on environment
func New(envType env.Type) Logger {
	switch envType {
	case env.TypeProd, env.TypeLocal:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		log.Logger = log.Logger.Level(zerolog.ErrorLevel)
		return &zeroLogger{
			logger: &log.Logger,
		}
	case env.TypeTest:
		return nopLogger()
	}
	panic("not handled")
}

func nopLogger() Logger {
	nopLogger := zerolog.Nop()
	return &zeroLogger{
		logger: &nopLogger,
	}
}
