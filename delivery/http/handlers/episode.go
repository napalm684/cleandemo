package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/napalm684/cleandemo/usecase"
)

//EpisodeHandler handles http requests for episode related REST endpoints
type EpisodeHandler struct {
	interactor usecase.EpisodeInteractor
}

//NewEpisodeHandler Constructor
func NewEpisodeHandler(interactor usecase.EpisodeInteractor) *EpisodeHandler {
	return &EpisodeHandler{
		interactor: interactor,
	}
}

//GetEpisodeByName handles requests for getting episodes by name
func (handler *EpisodeHandler) GetEpisodeByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/episodes/")
	episode, err := handler.interactor.GetEpisodeByName(name)

	if err != nil {
		http.Error(w, "Error retrieving episode details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episode)
}
