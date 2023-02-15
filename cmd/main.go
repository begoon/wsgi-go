package main

import (
	"os"
	"strconv"

	"github.com/begoon/wsgi-go/pkg/client"
	"github.com/begoon/wsgi-go/pkg/wsgi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("* zerolog initialized")

	serve := len(os.Args) == 1

	if serve {
		wsgi.Serve()
	} else {
		client.Client()
	}
}
