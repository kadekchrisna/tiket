package domains

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"tiket.vip/src/infrastructures/configs"
)

func (Event) TableName() string {
	return "event"
}

type Event struct {
	ID         string         `gorm:"column:id_event;PRIMARY_KEY" json:"id_event" `
	IDLocation string         `gorm:"column:id_location" json:"id_location" validate:"required"`
	Location   Location       `gorm:"foreignkey:IDLocation" json:"location"`
	Name       string         `json:"name" validate:"required"`
	Desc       string         `json:"desc" validate:"required"`
	StartDate  string         `gorm:"column:start_date" json:"start_date" validate:"required"`
	EndDate    string         `gorm:"column:end_date" json:"end_date" validate:"required"`
	CreatedAt  sql.NullString `json:"created_at"`
	UpdatedAt  sql.NullString `json:"updated_at"`
}

type EventPagi struct {
	EventName      string `json:"event_name" `
	EventCreatedAt string `json:"event_created_at" `
	LocName        string `json:"loc_name" `
	StartDate      string `json:"start_date" `
	EndDate        string `json:"end_date" `
	LocAddress     string `json:"loc_address" `
	LocCountry     string `json:"loc_country" `
	Order          string `json:"order" `
	Limit          int    `json:"limit" `
	Offset         int    `json:"offset" `
}

type Events struct {
	Events []Event `json:"events" `
	Total  int64   `json:"total" `
}

type EventRepo interface {
	GetAllEvents() (*[]Event, error)
	GetAllEventsPaginate(EventPagi) (*Events, error)
	GetEvent(string) (*Event, error)
	CreateEvent(Event) (*Event, error)
	UpdateEvent(Event) (*Event, error)
	DeleteEvent(string) (*string, error)
}

type EventUseCase interface {
	GetAllEvents() (*configs.ResponseSuccess, *configs.ResponseError)
	GetAllEventsPaginate(EventPagi) (*configs.ResponseSuccess, *configs.ResponseError)
	GetEvent(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	DeleteEvent(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	CreateEvent(Event) (*configs.ResponseSuccess, *configs.ResponseError)
	UpdateEvent(Event) (*configs.ResponseSuccess, *configs.ResponseError)
}

type EventController interface {
	GetAllEvents(echo.Context) error
	GetAllEventsPaginate(echo.Context) error
	GetEvent(echo.Context) error
	DeleteEvent(echo.Context) error
	CreateEvent(echo.Context) error
	UpdateEvent(echo.Context) error
}

func (e *EventPagi) Validate(limit string, offset string, order string) error {

	if valid.IsNotNull(e.EventCreatedAt) {
		cd, errCd := time.Parse(
			"2006-01-02 15:04:05",
			e.EventCreatedAt)

		if errCd != nil {
			return errors.New("Created date must be date time format")
		}

		e.EventCreatedAt = cd.Format("2006-01-02 15:04:05")

	}

	if valid.IsNotNull(e.StartDate) && valid.IsNotNull(e.EndDate) {
		sd, errSd := time.Parse(
			"2006-01-02 15:04:05",
			e.StartDate)

		ed, errEd := time.Parse(
			"2006-01-02 15:04:05",
			e.EndDate)

		if errSd != nil || errEd != nil {
			return errors.New("Created date, start date, and end date must be date time format")
		}

		if sd.After(ed) || sd == ed {
			return errors.New("Start date must be before end date")
		}
		e.EndDate = ed.Format("2006-01-02 15:04:05")
		e.StartDate = sd.Format("2006-01-02 15:04:05")
	}

	e.Order = "event.created_at desc"

	if valid.IsNotNull(order) {
		if len(strings.Split(order, " ")) > 1 && (strings.Contains(" asc", order) || strings.Contains(" desc", order)) {
			e.Order = order
		}
	}

	if valid.IsNull(limit) {
		e.Limit = 5
	} else {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return err
		}
		if e.Limit < 0 {
			e.Limit = 5
		}
		e.Limit = l
	}

	if valid.IsNull(offset) {
		e.Offset = 0
	} else {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return err
		}
		if e.Offset <= 0 {
			e.Offset = 0
		}
		e.Offset = l
	}
	return nil
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

type DB interface {
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *DB)
}
