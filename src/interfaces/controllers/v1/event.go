package controllers

import (
	"github.com/labstack/echo/v4"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type EventController struct {
	UseCase domains.EventUseCase
}

func NewEventController(uc domains.EventUseCase) domains.EventController {
	return &EventController{
		UseCase: uc,
	}
}

func (ec *EventController) GetAllEvents(c echo.Context) error {
	result, err := ec.UseCase.GetAllEvents()
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}
func (ec *EventController) GetAllEventsPaginate(c echo.Context) error {
	ep := domains.EventPagi{
		EventName:      c.QueryParam("eventName"),
		EventCreatedAt: c.QueryParam("eventCreated"),
		LocName:        c.QueryParam("locName"),
		LocAddress:     c.QueryParam("locAddress"),
		LocCountry:     c.QueryParam("locCountry"),
		EndDate:        c.QueryParam("endDate"),
		StartDate:      c.QueryParam("startDate"),
	}

	if err := ep.Validate(c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("order")); err != nil {
		rf := configs.ResponseError{
			Status:  400,
			Message: "FAILED",
			Cause:   err.Error(),
		}
		return c.JSON(rf.Status, rf)
	}
	result, err := ec.UseCase.GetAllEventsPaginate(ep)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (ec *EventController) GetEvent(c echo.Context) error {
	result, err := ec.UseCase.GetEvent(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (ec *EventController) CreateEvent(c echo.Context) error {
	var event domains.Event
	if err := event.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := ec.UseCase.CreateEvent(event)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (ec *EventController) UpdateEvent(c echo.Context) error {
	var event domains.Event
	if err := event.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := ec.UseCase.UpdateEvent(event)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (ec *EventController) DeleteEvent(c echo.Context) error {
	result, err := ec.UseCase.DeleteEvent(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}
