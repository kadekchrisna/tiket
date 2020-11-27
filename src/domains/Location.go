package domains

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"tiket.vip/src/infrastructures/configs"
)

func (Location) TableName() string {
	return "location"
}

type Location struct {
	ID        string  `gorm:"column:id_location;PRIMARY_KEY" json:"id_location"`
	Name      string  `json:"name" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Street    string  `json:"street" validate:"required"`
	City      string  `json:"city" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Zip       string  `json:"zip" validate:"required"`
	Latitude  string  `json:"latitude" validate:"required"`
	Longitude string  `json:"longitude" validate:"required"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type LocRepo interface {
	GetLocation(string) (*Location, error)
	CreateLocation(Location) (*Location, error)
}
type LocUseCase interface {
	CreateLocation(Location) (*configs.ResponseSuccess, *configs.ResponseError)
	SearchItem(query EsQuery) (*configs.ResponseSuccess, *configs.ResponseError)
}

type LocController interface {
	CreateLocation(echo.Context) error
	SearchItem(echo.Context) error
}

func (l *Location) Bind(c echo.Context) error {
	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}
	return nil
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	var location Location
	var count int64
	if l.ID == "" {
		l.ID = uuid.New().String()
	}
	for {
		if err := tx.Model(&location).Where("id_location = ?", l.ID).Count(&count).Error; err != nil {
			return err
		}
		if count < 1 {
			break
		}
		l.ID = uuid.New().String()
	}

	return
}
