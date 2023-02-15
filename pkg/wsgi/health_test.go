package wsgi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/health")

	if assert.NoError(t, HealthHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actual Health
		if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, actual, Health{Version: Version})
	}
}
