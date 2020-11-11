package middlewares_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiket.vip/src/infrastructures/middlewares"
)

func OpenDBConnection(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, errMock := sqlmock.New()
	assert.NoError(t, errMock)

	DB, errDb := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, errDb)

	return dbMock, DB, mock
}

func TestMiddleAuth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, DB, _ := OpenDBConnection(t)
	authMid := middlewares.NewAuthMiddleware(DB)

	authMid.MiddleGate()(echo.NotFoundHandler)(c)
	assert.Equal(t, "up-to-the-moon", rec.Header().Get("X-ORIGIN"))
}
