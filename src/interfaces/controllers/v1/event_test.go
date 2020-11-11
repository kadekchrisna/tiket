package controllers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"tiket.vip/src/domains"
	"tiket.vip/src/domains/mocks"
	"tiket.vip/src/infrastructures/configs"
	"tiket.vip/src/infrastructures/middlewares"
	"tiket.vip/src/interfaces/controllers/v1"
)

func TestGetEventSuccess(t *testing.T) {
	var mockEvent domains.Event
	errFakeEvent := faker.FakeData(&mockEvent)
	assert.NoError(t, errFakeEvent)

	mockSuccess := configs.Success(200, "OK", mockEvent)
	mockEvUseCase := new(mocks.EventUseCase)
	mockEvUseCase.On("GetEvent", mock.Anything).Return(mockSuccess, nil)

	eventC := controllers.NewEventController(mockEvUseCase)

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/event/get_info/%s", mockEvent.ID), strings.NewReader(""))
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	errH := eventC.GetEvent(c)
	require.NoError(t, errH)
}

func TestGetEventFailed(t *testing.T) {
	var mockEvent domains.Event
	errFakeEvent := faker.FakeData(&mockEvent)
	assert.NoError(t, errFakeEvent)

	mockFailed := configs.Failed(400, "FAIL", errors.New("ERROR").Error())
	mockEvUseCase := new(mocks.EventUseCase)
	mockEvUseCase.On("GetEvent", mock.Anything).Return(nil, mockFailed)

	eventC := controllers.NewEventController(mockEvUseCase)

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/event/get_info/%s", mockEvent.ID), strings.NewReader(""))
	rec := httptest.NewRecorder()

	e := echo.New()
	errorMidd := middlewares.NewErrorMiddleware()
	e.HTTPErrorHandler = errorMidd.HttpErrorHandler
	c := e.NewContext(req, rec)

	errH := eventC.GetEvent(c)
	assert.NoError(t, errH)
	// fmt.Println(rec)
	// assert.Equal(t, mockFailed.Cause, errH.Error())

	var failed configs.ResponseError
	errUn := json.Unmarshal(rec.Body.Bytes(), &failed)
	assert.NoError(t, errUn)

	assert.Equal(t, mockFailed.Cause, failed.Cause)

}
