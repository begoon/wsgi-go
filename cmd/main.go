package main

import (
	"os"

	"github.com/begoon/wsgi-go/wsgi"
)

func main() {
	serve := len(os.Args) > 1 && os.Args[1] == "serve"

	if serve {
		wsgi.Serve()
	} else {
		Connect()
	}
}
