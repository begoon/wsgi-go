package wsgi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	healthJSON = `{"version":"0.0.0"}` + "\n"
)

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")

	if assert.NoError(t, HealthHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, healthJSON, rec.Body.String())
		var actual Health
		if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, actual, Health{Version: "0.0.0"})
	}
}
