// Package helper common used settings functions
package helper

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/sirupsen/logrus"
)

func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return logger
}

func Measure(tn time.Time) {
	logrus.Println("took", time.Since(tn))
}
