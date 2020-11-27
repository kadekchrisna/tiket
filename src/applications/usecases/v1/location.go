package usecase

import (
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/configs"
)

type LocationUseCase struct {
	LocRepo  domains.LocRepo
	ItemRepo domains.ItemRepo
}

func NewLocUseCase(lr domains.LocRepo, ir domains.ItemRepo) domains.LocUseCase {
	return &LocationUseCase{
		LocRepo:  lr,
		ItemRepo: ir,
	}
}

func (l *LocationUseCase) CreateLocation(loc domains.Location) (*configs.ResponseSuccess, *configs.ResponseError) {
	location, errLoc := l.LocRepo.CreateLocation(loc)
	if errLoc != nil {
		return nil, configs.Failed(400, "FAILED", errLoc.Error())
	}

	return configs.Success(200, "SUCCESS", location), nil
}

func (l *LocationUseCase) SearchItem(query domains.EsQuery) (*configs.ResponseSuccess, *configs.ResponseError) {
	items, err := l.ItemRepo.Search(query)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "SUCCESS", items), nil
}
