package main

import (
	"log"
	"net/http"

	"github.com/yudhapratama10/search-service/infrastructures/elasticsearch"
)

func main() {
	log.Println("Starting Search Service")

	_, err := elasticsearch.GetClient()
	if err != nil {
		return
	}

	log.Println("Elastic Search Client Started")

	http.HandleFunc("/search", handler)

}
