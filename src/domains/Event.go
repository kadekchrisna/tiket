package domains

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"tiket.vip/src/infrastructures/configs"
)

func (Event) TableName() string {
	return "event"
}

type Event struct {
	ID         string    `gorm:"column:id_event;PRIMARY_KEY" json:"id_event" `
	IDLocation string    `gorm:"column:id_location" json:"id_location" validate:"required"`
	Location   Location  `gorm:"foreignkey:IDLocation" json:"location"`
	Name       string    `json:"name" validate:"required"`
	Desc       string    `json:"desc" validate:"required"`
	StartDate  string    `gorm:"column:start_date" json:"start_date" validate:"required"`
	EndDate    string    `gorm:"column:end_date" json:"end_date" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EventRepo interface {
	GetAllEvents() (*[]Event, error)
	GetEvent(string) (*Event, error)
	CreateEvent(Event) (*Event, error)
	UpdateEvent(Event) (*Event, error)
	DeleteEvent(string) (*string, error)
	GetLocation(string) (*Location, error)
}

type EventUseCase interface {
	GetAllEvents() (*configs.ResponseSuccess, *configs.ResponseError)
	GetEvent(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	DeleteEvent(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	CreateEvent(Event) (*configs.ResponseSuccess, *configs.ResponseError)
	UpdateEvent(Event) (*configs.ResponseSuccess, *configs.ResponseError)
}

type EventController interface {
	GetAllEvents(echo.Context) error
	GetEvent(echo.Context) error
	DeleteEvent(echo.Context) error
	CreateEvent(echo.Context) error
	UpdateEvent(echo.Context) error
}

func (e *Event) Bind(c echo.Context) error {
	if err := c.Bind(e); err != nil {
		return err
	}

	if err := c.Validate(e); err != nil {
		return err
	}

	sd, errSd := time.Parse(
		"2006-01-02 15:04:05",
		e.StartDate)
	ed, errEd := time.Parse(
		"2006-01-02 15:04:05",
		e.EndDate)

	if errSd != nil || errEd != nil {
		return errors.New("Start date and end date must be date time format")
	}

	if sd.After(ed) || sd == ed {
		return errors.New("Start date must be before end date")
	}
	return nil
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	var event Event
	var count int64
	for {
		e.ID = uuid.New().String()
		if err := tx.Model(&event).Where("id_event = ?", e.ID).Count(&count).Error; err != nil {
			return err
		}
		if count < 1 {
			break
		}
	}

	return
}
