package handler

import (
	"net/http"

	usecase "github.com/yudhapratama10/search-service/usecases"
)

type FootballHandler struct {
	footballUsecase usecase.FootballUsecaseContract
}

type FootballHandlerContract interface {
	SearchFootballClub(w http.ResponseWriter, r *http.Request)
	Autocomplete(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(footballUsecase usecase.FootballUsecaseContract) FootballHandlerContract {
	return &FootballHandler{
		footballUsecase: footballUsecase,
	}
}
