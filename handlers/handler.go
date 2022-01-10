package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yudhapratama10/search-service/model"
)

func (handler *FootballHandler) SearchFootballClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var req model.SearchParam

		// fmt.Println(r.URL.Query())

		// err := json.NewDecoder().Decode(&req)

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		take, _ := strconv.Atoi(r.URL.Query().Get("take"))
		hasStadium, _ := strconv.ParseBool(r.URL.Query().Get("hasstadium"))

		req = model.SearchParam{
			Keyword:    r.URL.Query().Get("keyword"),
			HasStadium: hasStadium,
			Page:       page,
			Take:       take,
		}

		cursorSearch, err := handler.footballUsecase.Search(req.Keyword, req.HasStadium, req.Page, req.Take)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(cursorSearch)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (handler *FootballHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var req model.SearchParam

		req = model.SearchParam{
			Keyword: r.URL.Query().Get("keyword"),
		}

		cursorAutocomplete, err := handler.footballUsecase.Autocomplete(req.Keyword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(cursorAutocomplete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
