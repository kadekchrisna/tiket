package datasources

import (
	"gorm.io/gorm"
	"tiket.vip/src/domains"
)

type Repos struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) domains.EventRepo {
	return &Repos{
		db: db,
	}
}
