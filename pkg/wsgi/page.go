package wsgi

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var tmpls *template.Template

func PageHandler(c echo.Context) error {
	page := c.Param("page")
	log.Info().Str("page", page)

	v := struct{ Version string }{Version}

	if tmpls == nil {
		tmpls = Must(template.ParseFS(templates, "templates/*"))
	}

	var b bytes.Buffer
	if err := tmpls.ExecuteTemplate(&b, page, &v); err != nil {
		panic(fmt.Sprintf("%v: parsing %s", err, page))
	}
	return c.String(http.StatusOK, b.String())
}
