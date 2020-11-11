package usecase

import (
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type LocationUseCase struct {
	LocRepo domains.LocRepo
}

func NewLocUseCase(lr domains.LocRepo) domains.LocUseCase {
	return &LocationUseCase{
		LocRepo: lr,
	}
}

func (l *LocationUseCase) CreateLocation(loc domains.Location) (*configs.ResponseSuccess, *configs.ResponseError) {
	location, errLoc := l.LocRepo.CreateLocation(loc)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	return configs.Success(200, "SUCCESS", location), nil
}
