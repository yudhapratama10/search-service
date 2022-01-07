package repository

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yudhapratama10/search-service/model"
)

type footballRepository struct {
	client *elasticsearch.Client
}

type FootballRepositoryContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	Synonym(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	Autocomplete(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
}

func NewRecipeRepository(client *elasticsearch.Client) FootballRepositoryContract {
	return &footballRepository{client: client}
}
