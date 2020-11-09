package datasources

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"tiket.vip/src/domains"
)

func (r *Repos) GetLocation(id string) (*domains.Location, error) {
	var loc domains.Location

	if err := r.db.Where("id_location = ?", id).First(&loc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fmt.Sprintf("location with id %s is not found", id))
		}
		return nil, err
	}
	return &loc, nil
}
