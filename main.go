package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/yudhapratama10/search-service/handlers"
	"github.com/yudhapratama10/search-service/infrastructures/elasticsearch"
	repository "github.com/yudhapratama10/search-service/repositories"
	usecase "github.com/yudhapratama10/search-service/usecases"
)

func main() {
	log.Println("Starting Search Service")

	esClient, err := elasticsearch.GetClient()
	if err != nil {
		return
	}

	log.Println("Elastic Search Client Started")

	repo := repository.NewRecipeRepository(esClient)
	uc := usecase.NewFootballUsecase(repo)
	handler := handler.NewProductHandler(uc)

	http.HandleFunc("/search", handler.SearchFootballClub)
	http.HandleFunc("/autocomplete", handler.Autocomplete)

	fmt.Println("Starting Service")
	http.ListenAndServe(":8080", nil)
}
