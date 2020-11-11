package datasources

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"tiket.vip/src/domains"
)

type LocRepos struct {
	db *gorm.DB
}

func NewLocRepo(db *gorm.DB) domains.LocRepo {
	return &LocRepos{
		db: db,
	}
}

func (r *LocRepos) GetLocation(id string) (*domains.Location, error) {
	var loc domains.Location
	if err := r.db.Where("id_location = ?", id).First(&loc).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New(fmt.Sprintf("location with id %s is not found", id))
		}
		return nil, err
	}
	return &loc, nil
}

func (r *LocRepos) CreateLocation(loc domains.Location) (*domains.Location, error) {
	var location domains.Location
	trx := r.db.Begin()
	result := trx.Create(&loc)
	if result.Error != nil {

		if err := trx.Rollback().Error; err != nil {
			return nil, err
		}
		return nil, result.Error
	}

	if err := trx.Commit().Error; err != nil {
		return nil, err
	}
	location = loc
	return &location, nil
}
