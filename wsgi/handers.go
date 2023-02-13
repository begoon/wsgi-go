package wsgi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Version string `json:"version"`
}

func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{Version: "0.0.0"})
}

type Process struct {
	N        int           `json:"n"`
	Duration time.Duration `json:"duration"`
}

func ProcessHandler(c echo.Context) error {
	n, _ := strconv.Atoi(c.Param("n"))
	duration, _ := strconv.Atoi(c.Param("duration"))

	p := Process{Duration: time.Duration(duration) * time.Millisecond, N: n}

	time.Sleep(p.Duration)
	return c.JSON(http.StatusOK, p)
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func AddUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Id = c.Response().Header().Get("X-Request-Id")
	return c.JSON(http.StatusCreated, u)
}
