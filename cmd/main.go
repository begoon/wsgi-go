package main

import (
	"os"
	"strconv"
	"time"

	"github.com/begoon/wsgi-go/pkg/client"
	"github.com/begoon/wsgi-go/pkg/wsgi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Version = "1.0.0"

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

	if os.Getenv("NICE") != "" {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	}

	log.Info().Msg("* zerolog initialized")

	serve := len(os.Args) == 1

	if serve {
		wsgi.Serve(Version)
	} else {
		client.Client()
	}
}
