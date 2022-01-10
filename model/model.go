package model

type FootballClub struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Tournaments []string `json:"tournaments"`
	Nation      string   `json:"nation"`
	HasStadium  bool     `json:"has_stadium"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
}

type SearchParam struct {
	Keyword    string `json:"keyword"`
	HasStadium bool   `json:"hasStadium"`
	Page       int    `json:"page"`
	Take       int    `json:"take"`
}

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
