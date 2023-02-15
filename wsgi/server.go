package wsgi

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/rs/zerolog/log"
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

func requestDumper(c echo.Context, reqBody, resBody []byte) {
	id := c.Response().Header()["X-Request-Id"][0]
	log.Info().
		Str("id", id).
		RawJSON("request", reqBody).
		RawJSON("response", bytes.TrimSpace(resBody)).
		Msgf("REQUEST/RESPONSE")
}

func requestDumperSkipper(c echo.Context) bool {
	u := c.Request().RequestURI
	return !strings.Contains(u, "/api") || strings.Contains(u, "/docs")
}

func Serve() {
	e := echo.New()

	r := echoswagger.New(e, "/api/docs", &echoswagger.Info{
		Title:   "Chiro API",
		Version: Version,
	})
	r.AddSecurityAPIKey("api_key", "", echoswagger.SecurityInHeader)

	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	printFS("root", root)

	e.StaticFS("/", echo.MustSubFS(root, "root"))

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: requestDumperSkipper,
		Handler: requestDumper,
	}))

	e.GET("/api", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/docs")
	})

	r.GET("/api/health", HealthHandler).
		AddResponse(http.StatusOK, "successful", Health{}, nil)

	r.GET("/api/process/:n/:duration", ProcessHandler).
		AddParamPath(int(0), "n", "process id").
		AddParamPath(int(0), "duration", "time in milliseconds").
		AddResponse(http.StatusOK, "successful", Process{}, nil)

	r.POST("/api/user", AddUserHandler).
		AddParamBody(User{}, "body", "User object", true).
		AddResponse(http.StatusCreated, "created successfully", nil, nil)

	r.GET("/api/ws/:cid", WebSocketHandler).
		AddParamPath(string(""), "cid", "client id")

	r.
		SetScheme("https", "http").
		SetResponseContentType("application/json").
		SetRequestContentType("application/json")

	var port string
	port, given := os.LookupEnv("PORT")
	if !given {
		port = "8000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
