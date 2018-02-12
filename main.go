package main

import (
	"net/http"

	"github.com/napalm684/cleandemo/delivery/http/handlers"
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	inmemRepo "github.com/napalm684/cleandemo/infrastructure/repository/inmemory"
	"github.com/napalm684/cleandemo/usecase"
)

func main() {
	repository := inmemRepo.NewEpisodeRepository(generateInMemoryData())
	episodeService := usecase.NewEpisodeService(repository)
	episodeHandler := handlers.NewEpisodeHandler(*episodeService)

	http.HandleFunc("/episodes/", episodeHandler.GetEpisodeByName)
	http.ListenAndServe(":8888", nil)
}

func generateInMemoryData() map[string]*episodeDomain.Episode {
	var m map[string]*episodeDomain.Episode
	m = make(map[string]*episodeDomain.Episode)
	m["Kitten"] = &episodeDomain.Episode{
		Name:    "Kitten",
		Season:  11,
		Episode: 6,
	}

	return m
}
