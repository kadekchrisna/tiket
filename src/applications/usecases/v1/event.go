package usecase

import (
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type EventUseCase struct {
	EvRepo  domains.EventRepo
	LocRepo domains.LocRepo
}

func NewEventUseCase(er domains.EventRepo, lr domains.LocRepo) domains.EventUseCase {
	return &EventUseCase{
		EvRepo:  er,
		LocRepo: lr,
	}
}

func (eu *EventUseCase) GetAllEvents() (*configs.ResponseSuccess, *configs.ResponseError) {
	events, err := eu.EvRepo.GetAllEvents()
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", events), nil
}

func (eu *EventUseCase) GetAllEventsPaginate(ep domains.EventPagi) (*configs.ResponseSuccess, *configs.ResponseError) {

	// logrus.WithFields(logrus.Fields{
	// 	"ep": ep,
	// }).Debug("ep")
	events, err := eu.EvRepo.GetAllEventsPaginate(ep)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", events), nil
}

func (eu *EventUseCase) GetEvent(id interface{}) (*configs.ResponseSuccess, *configs.ResponseError) {
	s, ok := id.(string)
	if !ok {
		return nil, configs.Failed(400, "FAILED", "id must be a string")
	}

	// t := time.Now()
	// fmt.Println(t.Format("2006-01-02 15:04:05"))
	events, err := eu.EvRepo.GetEvent(s)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", events), nil
}

func (eu *EventUseCase) CreateEvent(ev domains.Event) (*configs.ResponseSuccess, *configs.ResponseError) {
	loc, errLoc := eu.LocRepo.GetLocation(ev.IDLocation)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	event, err := eu.EvRepo.CreateEvent(ev)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	event.Location = *loc
	return configs.Success(200, "OK", event), nil
}

func (eu *EventUseCase) UpdateEvent(ev domains.Event) (*configs.ResponseSuccess, *configs.ResponseError) {
	loc, errLoc := eu.LocRepo.GetLocation(ev.IDLocation)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	event, err := eu.EvRepo.UpdateEvent(ev)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	event.Location = *loc
	return configs.Success(200, "OK", event), nil
}

func (eu *EventUseCase) DeleteEvent(id interface{}) (*configs.ResponseSuccess, *configs.ResponseError) {
	s, ok := id.(string)
	if !ok {
		return nil, configs.Failed(400, "FAILED", "id must be a string")
	}
	event, errCheck := eu.EvRepo.GetEvent(s)
	if errCheck != nil {
		return nil, configs.Failed(400, "FAILED", errCheck.Error())
	}

	id, errDel := eu.EvRepo.DeleteEvent(event.ID)
	if errDel != nil {
		return nil, configs.Failed(400, "FAILED", errDel.Error())
	}
	return configs.Success(200, "OK", event), nil
}
