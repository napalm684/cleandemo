package main

import (
	"net/http"

	"github.com/napalm684/cleandemo/delivery/http/handlers"
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/napalm684/cleandemo/usecase"
	repositoryMock "github.com/napalm684/cleandemo/usecase/repository/mocks"
)

func main() {
	repository := &repositoryMock.Episode{} //TODO: Temporary for testing
	repository.On(
		"GetEpisodeByName",
		"Kitten",
	).Return(&episodeDomain.Episode{
		Name:    "Kitten",
		Season:  11,
		Episode: 6,
	}, nil)

	episodeInteractor := usecase.NewEpisodeInteractor(repository)
	episodeHandler := handlers.NewEpisodeHandler(*episodeInteractor)

	http.HandleFunc("/episodes/", episodeHandler.GetEpisodeByName)
	http.ListenAndServe(":8888", nil)
}
