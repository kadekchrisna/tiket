package datasources

import (
	"errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiket.vip/src/domains"
)

type EventRepos struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) domains.EventRepo {
	return &EventRepos{
		db: db,
	}
}

func (r *EventRepos) GetAllEvents() (*[]domains.Event, error) {
	var events []domains.Event
	if err := r.db.Preload("Location").Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}

func (r *EventRepos) GetEvent(id string) (*domains.Event, error) {
	var event domains.Event
	var err error
	if err = r.db.Preload("Location").Where("id_event = ?", id).First(&event).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New(fmt.Sprintf("event with id %s is not found", id))
		}
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"id": err,
	}).Debug("id------")
	return &event, nil
}

func (r *EventRepos) CreateEvent(ev domains.Event) (*domains.Event, error) {
	var event domains.Event
	result := r.db.Omit(clause.Associations).Create(&ev)

	if result.Error != nil {
		return nil, result.Error
	}
	event = ev
	logrus.WithFields(logrus.Fields{
		"ev": ev,
	}).Debug("ev-----")
	return &event, nil
}

func (r *EventRepos) UpdateEvent(ev domains.Event) (*domains.Event, error) {
	var event domains.Event
	result := r.db.Model(&event).Where("id_event = ?", ev.ID).Updates(ev)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New(fmt.Sprintf("event with id %s is not found", ev.ID))
	}
	event = ev
	return &event, nil
}

func (r *EventRepos) DeleteEvent(id string) (*string, error) {
	if err := r.db.Where("id_event = ?", id).Delete(&domains.Event{}).Error; err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *EventRepos) GetAllEventsPaginateCount(trx *gorm.DB) (*int64, error) {
	var count int64
	if err := trx.Count(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}

func (r *EventRepos) GetAllEventsPaginateQuery(ep domains.EventPagi) *gorm.DB {
	trx := r.db.Table("event event").
		Select("event.id_event, event.id_location, event.name, event.desc, event.start_date, event.end_date, event.created_at, event.updated_at, location.id_location, location.name, location.address, location.street, location.city, location.country, location.zip, location.latitude, location.longitude, location.created_at, location.updated_at").
		Joins("join location location on event.id_location = location.id_location")

	if valid.IsNotNull(ep.EventName) {
		trx = trx.Where("LOWER(event.name) LIKE ?", "%"+strings.ToLower(ep.EventName)+"%")
	}

	if valid.IsNotNull(ep.LocName) {
		trx = trx.Where("LOWER(location.name) LIKE ?", "%"+strings.ToLower(ep.LocName)+"%")
	}
	if valid.IsNotNull(ep.LocAddress) {
		trx = trx.Where("LOWER(location.address) LIKE ?", "%"+strings.ToLower(ep.EventName)+"%")
	}

	if valid.IsNotNull(ep.LocName) {
		trx = trx.Where("LOWER(location.name) LIKE ?", "%"+strings.ToLower(ep.LocName)+"%")
	}
	return trx
}
func (r *EventRepos) GetAllEventsPaginate(ep domains.EventPagi) (*domains.Events, error) {
	trx := r.GetAllEventsPaginateQuery(ep)
	count, errCount := r.GetAllEventsPaginateCount(trx)
	if errCount != nil {
		return nil, errCount
	}

	var events domains.Events
	if *count > 0 {
		events.Total = *count
		trx.Limit(ep.Limit).Offset(ep.Offset).Order(ep.Order)
		rows, errRows := trx.Rows()
		if errRows != nil {
			return nil, errRows
		}
		defer rows.Close()
		for rows.Next() {
			var event domains.Event
			if err := rows.Scan(&event.ID, &event.IDLocation, &event.Name, &event.Desc, &event.StartDate, &event.EndDate, &event.CreatedAt, &event.UpdatedAt, &event.Location.ID, &event.Location.Name, &event.Location.Address, &event.Location.Street, &event.Location.City, &event.Location.Country, &event.Location.Zip, &event.Location.Latitude, &event.Location.Longitude, &event.Location.CreatedAt, &event.Location.UpdatedAt); err != nil {
				return nil, err
			}
			events.Events = append(events.Events, event)
		}
	} else {
		events.Total = 0
	}

	return &events, nil
}
