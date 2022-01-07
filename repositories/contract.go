package repository

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/yudhapratama10/search-service/model"
)

type SearchResult struct {
	Hits HitsSearchResult `json:"hits"`
}

type HitsSearchResult struct {
	Total TotalResult  `json:"total"`
	Hits  []HitsResult `json:"hits"`
}

type TotalResult struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type HitsResult struct {
	Id     string                 `json:"_id"`
	Score  float32                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

// type SourceResult struct {
// 	Source `json:"_source"`
// }

type footballRepository struct {
	client *elasticsearch.Client
}

type FootballRepositoryContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	// Synonym(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	// Autocomplete(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
}

func NewRecipeRepository(client *elasticsearch.Client) FootballRepositoryContract {
	return &footballRepository{client: client}
}
