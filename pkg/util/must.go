package util

import (
	"github.com/rs/zerolog/log"
)

func Must[T any](r T, err error) T {
	if err != nil {
		log.Fatal().Err(err).Msg("Must")
	}
	return r
}
