package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

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
