package main

import (
	"os"

	"github.com/begoon/wsgi-go/pkg/client"
	"github.com/begoon/wsgi-go/pkg/wsgi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("* zerolog initialized")

	serve := len(os.Args) == 1

	if serve {
		wsgi.Serve()
	} else {
		client.Client()
	}
}
