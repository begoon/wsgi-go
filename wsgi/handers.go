package wsgi

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Version string `json:"version"`
}

func HealthHandler(c echo.Context) error {
	h := Health{Version: "0.0.0"}
	if err := c.Bind(h); err != nil {
		return err
	}
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return c.JSON(http.StatusOK, h)
}

type Process struct {
	N        int           `json:"n"`
	Duration time.Duration `json:"duration"`
}

func ProcessHandler(c echo.Context) error {
	n, _ := strconv.Atoi(c.Param("n"))
	duration, _ := strconv.Atoi(c.Param("duration"))

	p := Process{Duration: time.Duration(duration) * time.Second, N: n}

	time.Sleep(p.Duration)
	return c.JSON(http.StatusOK, p)
}
