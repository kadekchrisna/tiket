package domains

import (
	"time"
)

func (Location) TableName() string {
	return "location"
}

type Location struct {
	ID        string    `gorm:"column:id_location;PRIMARY_KEY" json:"id_location"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Zip       string    `json:"zip"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LocRepo interface {
	GetLocation(string) (*Location, error)
}
