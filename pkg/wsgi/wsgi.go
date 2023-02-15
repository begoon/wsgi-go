package wsgi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"embed"

	"github.com/begoon/wsgi-go/pkg/util"
	"github.com/begoon/wsgi-go/pkg/wsgi/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/rs/zerolog/log"
)

//go:embed root/*
var root embed.FS

func requestDumper(c echo.Context, reqBody, resBody []byte) {
	id := c.Response().Header()["X-Request-Id"][0]
	req := map[string]interface{}{}
	if len(reqBody) > 0 {
		if err := json.Unmarshal(reqBody, &req); err != nil {
			panic(err)
		}
	}
	log.Info().
		Str("id", id).
		Interface("request", req).
		RawJSON("response", bytes.TrimSpace(resBody)).
		Msgf("REQUEST/RESPONSE")
}

func requestDumperSkipper(c echo.Context) bool {
	u := c.Request().RequestURI
	return !strings.Contains(u, "/api") || strings.Contains(u, "/docs")
}

var opened = []string{"/health", "/process"}

const apiKey = "secret"

func apiKeyValidator(key string, c echo.Context) (bool, error) {
	return key == apiKey, nil
}

func apiKeySkipper(c echo.Context) bool {
	u := c.Request().URL.RequestURI()
	for _, e := range opened {
		if strings.Contains(u, e) {
			return true
		}
	}
	return false
}

func Serve(version string) {
	handler.Version = version

	e := echo.New()

	r := echoswagger.New(e, "/api/docs", &echoswagger.Info{
		Title:   "Chiro API",
		Version: version,
	}).
		SetScheme("https", "http").
		SetResponseContentType("application/json").
		SetRequestContentType("application/json")

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-Api-Key",
		Validator: apiKeyValidator,
		Skipper:   apiKeySkipper,
	}))

	r.AddSecurityAPIKey("X-Api-Key", "", echoswagger.SecurityInHeader)

	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	util.PrintFS("root", root)

	e.StaticFS("/", echo.MustSubFS(root, "root"))

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: requestDumperSkipper,
		Handler: requestDumper,
	}))

	e.GET("/api", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/docs")
	})

	r.GET("/api/health", handler.HealthHandler).
		AddResponse(http.StatusOK, "successful", handler.Health{}, nil)

	r.GET("/api/process/:n/:duration", handler.ProcessHandler).
		AddParamPath(int(0), "n", "process id").
		AddParamPath(int(0), "duration", "time in milliseconds").
		AddResponse(http.StatusOK, "successful", handler.Process{}, nil)

	r.POST("/api/user", handler.AddUserHandler).
		AddParamBody(handler.User{}, "body", "User object", true).
		AddResponse(http.StatusCreated, "created successfully", nil, nil)

	r.GET("/api/ws/:cid", handler.WebSocketHandler).
		AddParamPath(string(""), "cid", "client id")

	r.GET("/api/panic", func(c echo.Context) error {
		panic("PANIC")
	})

	r.GET("/page/:page", handler.PageHandler).
		AddParamPath(string(""), "page", "page path").
		AddResponse(http.StatusOK, "successful", "", nil)

	var port string
	port, given := os.LookupEnv("PORT")
	if !given {
		port = "8000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
