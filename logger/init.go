package logger

import (
	"os"

	"github.com/blueberry-guy/go-kit/env"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(envType env.Type) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	switch envType {
	case env.TypeProd:
		log.Logger = log.Logger.Level(zerolog.ErrorLevel)
	}
	zl := &zeroLogger{}
	zl.logger = &log.Logger
	return zl
}
