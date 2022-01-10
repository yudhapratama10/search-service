package usecase

import (
	"github.com/yudhapratama10/search-service/model"
)

func (usecase *footballUsecase) Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {

	data, err := usecase.repo.Search(keyword, hasStadium, page, take)
	if err != nil {
		return []model.FootballClub{}, err
	}

	return data, nil
}

func (usecase *footballUsecase) Autocomplete(keyword string) ([]model.FootballClub, error) {

	data, err := usecase.repo.Autocomplete(keyword)
	if err != nil {
		return []model.FootballClub{}, err
	}

	return data, nil
}
