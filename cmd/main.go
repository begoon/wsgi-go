package main

import (
	"os"

	"github.com/begoon/wsgi-go/wsgi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("* zerolog initialized")

	serve := len(os.Args) > 1 && os.Args[1] == "serve"

	if serve {
		wsgi.Serve()
	} else {
		Connect()
	}
}
