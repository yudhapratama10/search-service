package usecase

import (
	"github.com/yudhapratama10/search-service/model"
	repository "github.com/yudhapratama10/search-service/repositories"
)

type footballUsecase struct {
	repo repository.FootballRepositoryContract
}

type FootballUsecaseContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
}

func NewFootballUsecase(repo repository.FootballRepositoryContract) FootballUsecaseContract {
	return &footballUsecase{repo: repo}
}
