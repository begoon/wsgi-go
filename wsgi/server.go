package wsgi

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
)

//go:embed root/*
var root embed.FS

func printFS(prefix string, fs embed.FS) {
	if p, err := fs.ReadDir(prefix); err == nil {
		for _, f := range p {
			info, err := f.Info()
			if err != nil {
				panic(fmt.Sprintf("unable to read info of %v", f))
			}
			path := prefix + "/" + f.Name()
			if info.IsDir() {
				printFS(path, fs)
			} else {
				fmt.Printf("%v (%d)\n", path, info.Size())
			}
		}
	} else {
		panic("embed files not found")
	}
}

func Serve() {
	e := echo.New()

	r := echoswagger.New(e, "/api/docs", nil)

	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	printFS("root", root)

	e.StaticFS("/", echo.MustSubFS(root, "root"))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		id := c.Response().Header()["X-Request-Id"][0]
		fmt.Printf("REQUEST  [%s] [%v]\n", id, string(reqBody))
		fmt.Printf("RESPONSE [%s] [%v]\n\n", id, strings.TrimSpace(string(resBody)))
	}))

	e.GET("/api", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/docs")
	})

	r.GET("/api/health", HealthHandler).
		AddParamBody(Health{}, "body", "User input struct", true).
		AddResponse(http.StatusOK, "successful", nil, nil)

	r.GET("/api/process/:n/:duration", ProcessHandler).
		AddParamPath(int(0), "n", "process id").
		AddParamPath(int(0), "duration", "time in milliseconds").
		AddResponse(http.StatusOK, "successful", Process{}, nil)

	r.POST("/api/user", AddUserHandler).
		AddParamBody(User{}, "body", "User object", true).
		AddResponse(http.StatusCreated, "successful", nil, nil)

	r.SetScheme("https", "http")
	r.SetResponseContentType("application/json")

	var port string
	port, given := os.LookupEnv("PORT")
	if !given {
		port = "8000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
