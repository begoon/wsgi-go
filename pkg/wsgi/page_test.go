package wsgi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("page")
	c.SetParamValues("index.html")

	if assert.NoError(t, PageHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, rec.Body.String(), "Root page, Version "+Version+"\n")
	}
}
