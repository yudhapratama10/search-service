package repository

import (
	"github.com/yudhapratama10/search-service/model"
)

func (repo *footballRepository) Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {
	return []model.FootballClub{}, nil
}

func (repo *footballRepository) Synonym(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {
	return []model.FootballClub{}, nil
}

func (repo *footballRepository) Autocomplete(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {
	return []model.FootballClub{}, nil
}
