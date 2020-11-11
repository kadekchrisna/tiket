package middlewares_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"tiket.vip/src/infrastructures/configs"
	"tiket.vip/src/infrastructures/middlewares"
)

func TestHttpErrorHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	errorMidd := middlewares.NewErrorMiddleware()

	expect := errors.New("ERROR")

	c := e.NewContext(req, rec)
	errorMidd.HttpErrorHandler(expect, c)

	var failed configs.ResponseError
	errUn := json.Unmarshal(rec.Body.Bytes(), &failed)
	assert.NoError(t, errUn)

	assert.Equal(t, expect.Error(), failed.Cause)
}
