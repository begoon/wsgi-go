package wsgi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" swagger:"required"`
}

func AddUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Id = c.Response().Header().Get("X-Request-Id")
	return c.JSON(http.StatusCreated, u)
}
