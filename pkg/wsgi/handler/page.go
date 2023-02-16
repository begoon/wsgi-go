package handler

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"

	"github.com/begoon/wsgi-go/pkg/util"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var tmpls *template.Template

//go:embed templates/*
var templates embed.FS

func PageHandler(c echo.Context) error {
	page := c.Param("page")
	log.Info().Str("page", page)

	v := struct{ Version string }{Version}

	if tmpls == nil {
		tmpls = util.Must(template.ParseFS(templates, "templates/*"))
	}

	var b bytes.Buffer
	if err := tmpls.ExecuteTemplate(&b, page, &v); err != nil {
		log.Error().Err(err).Msgf("parsing %s", page)
		return err
	}
	return c.String(http.StatusOK, b.String())
}
