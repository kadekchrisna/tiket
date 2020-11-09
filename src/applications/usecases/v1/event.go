package usecase

import (
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type EventUseCase struct {
	Repo domains.EventRepo
}

func NewEventUseCase(er domains.EventRepo) domains.EventUseCase {
	return &EventUseCase{
		Repo: er,
	}
}

func (eu *EventUseCase) GetAllEvents() (*configs.ResponseSuccess, *configs.ResponseError) {
	events, err := eu.Repo.GetAllEvents()
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
	events, err := eu.Repo.GetEvent(s)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", events), nil
}

func (eu *EventUseCase) CreateEvent(ev domains.Event) (*configs.ResponseSuccess, *configs.ResponseError) {
	loc, errLoc := eu.Repo.GetLocation(ev.IDLocation)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	event, err := eu.Repo.CreateEvent(ev)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	event.Location = *loc
	return configs.Success(200, "OK", event), nil
}

func (eu *EventUseCase) UpdateEvent(ev domains.Event) (*configs.ResponseSuccess, *configs.ResponseError) {
	loc, errLoc := eu.Repo.GetLocation(ev.IDLocation)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	event, err := eu.Repo.UpdateEvent(ev)
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
	event, errCheck := eu.Repo.GetEvent(s)
	if errCheck != nil {
		return nil, configs.Failed(400, "FAILED", errCheck.Error())
	}

	id, errDel := eu.Repo.DeleteEvent(event.ID)
	if errDel != nil {
		return nil, configs.Failed(400, "FAILED", errDel.Error())
	}
	return configs.Success(200, "OK", event), nil
}
