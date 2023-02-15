package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Version string `json:"version"`
}

func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{Version: Version})
}
