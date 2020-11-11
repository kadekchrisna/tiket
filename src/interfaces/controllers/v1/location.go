package controllers

import (
	"github.com/labstack/echo/v4"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type LocationController struct {
	UseCase domains.LocUseCase
}

func NewLocController(uc domains.LocUseCase) domains.LocController {
	return &LocationController{
		UseCase: uc,
	}
}

func (lc *LocationController) CreateLocation(c echo.Context) error {
	var location domains.Location
	if err := location.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := lc.UseCase.CreateLocation(location)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}
