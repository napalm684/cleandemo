package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
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
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episode)
}

func generateDomainEpisode(name string, season int, episode int) *episodeDomain.Episode {
	return &episodeDomain.Episode{
		Name:    name,
		Season:  season,
		Episode: episode,
	}
}
