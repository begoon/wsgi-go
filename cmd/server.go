package main

import (
	"fmt"

	"github.com/begoon/wsgi-go/wsgi"
)

func main() {
	fmt.Println("@ console")

	wsgi.Serve()
}
