package wsgi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
)

func Serve() {

	e := echo.New()

	r := echoswagger.New(e, "/docs", nil)

	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		id := c.Response().Header()["X-Request-Id"][0]
		fmt.Printf("REQUEST  [%s] [%v]\n", id, string(reqBody))
		fmt.Printf("RESPONSE [%s] [%v]\n\n", id, strings.TrimSpace(string(resBody)))
	}))

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/docs")
	})

	r.GET("/health", HealthHandler).
		AddParamBody(Health{}, "body", "User input struct", true).
		AddResponse(http.StatusCreated, "successful", nil, nil)

	r.GET("/process/:n/:duration", ProcessHandler).
		AddParamPath(int(0), "n", "process id").
		AddParamPath(int(0), "duration", "time in milliseconds").
		AddResponse(http.StatusOK, "successful", Process{}, nil)

	r.SetScheme("https", "http")
	r.SetResponseContentType("application/json")

	e.Logger.Fatal(e.Start(":9999"))
}
