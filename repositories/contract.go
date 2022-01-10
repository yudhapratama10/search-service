package repository

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/yudhapratama10/search-service/model"
)

// type SourceResult struct {
// 	Source `json:"_source"`
// }

type footballRepository struct {
	client *elasticsearch.Client
}

type FootballRepositoryContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	// Synonym(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	Autocomplete(keyword string) ([]model.FootballClub, error)
}

func NewRecipeRepository(client *elasticsearch.Client) FootballRepositoryContract {
	return &footballRepository{client: client}
}
