package handler

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" swagger:"required"`
	Hash string `json:"hash" swagger:"omitempty"`
	Ulid string `json:"ulid" swagger:"omitempty"`
}

func AddUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Id = c.Response().Header().Get("X-Request-Id")
	u.Hash = base64.StdEncoding.EncodeToString([]byte("RAW"))
	u.Ulid = ulid.Make().String()
	return c.JSON(http.StatusCreated, u)
}
