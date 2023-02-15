package wsgi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	sent := User{Name: "-name-", Id: "-id-"}
	raw, err := json.Marshal(sent)
	assert.Nil(t, err, "error marshalling")

	reader := bytes.NewReader(raw)
	req := httptest.NewRequest(http.MethodPost, "/api/user", reader)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := echo.New().NewContext(req, rec)
	c.SetPath("/api/user")
	c.Response().Header().Add("X-Request-Id", "-x-request-id-")

	if assert.NoError(t, AddUserHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actual User
		if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, actual, User{Name: "-name-", Id: "-x-request-id-"})
	}
}