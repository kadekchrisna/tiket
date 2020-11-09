package datasources

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"tiket.vip/src/domains"
)

func (r *Repos) GetAllEvents() (*[]domains.Event, error) {
	var events []domains.Event
	if err := r.db.Preload("Location").Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}

func (r *Repos) GetEvent(id string) (*domains.Event, error) {
	var event domains.Event
	var err error
	if err = r.db.Preload("Location").Where("id_event = ?", id).First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, errors.New(fmt.Sprintf("event with id %s is not found", id))
		}
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"id": err,
	}).Debug("id------")
	return &event, nil
}

func (r *Repos) CreateEvent(ev domains.Event) (*domains.Event, error) {
	var event domains.Event
	result := r.db.Create(&ev)

	if result.Error != nil {
		return nil, result.Error
	}
	event = ev
	// logrus.WithFields(logrus.Fields{
	// 	"ev": ev,
	// }).Debug("ev")
	return &event, nil
}

func (r *Repos) UpdateEvent(ev domains.Event) (*domains.Event, error) {
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

func (r *Repos) DeleteEvent(id string) (*string, error) {
	if err := r.db.Where("id_event = ?", id).Delete(&domains.Event{}).Error; err != nil {
		return nil, err
	}
	return &id, nil
}
