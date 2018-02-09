package handlers_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/napalm684/cleandemo/delivery/http/handlers"
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/napalm684/cleandemo/usecase"
	repositoryMock "github.com/napalm684/cleandemo/usecase/repository/mocks"
)

func TestEpisodeHandlerGetEpisodeByName(t *testing.T) {
	// Arrange
	t.Parallel()
	episode := generateDomainEpisode("Kitten", 11, 6)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/episodes/"+episode.Name, nil)
	if err != nil {
		t.Fatal(err)
	}

	repository := &repositoryMock.Episode{}
	repository.On(
		"GetEpisodeByName",
		episode.Name,
	).Return(episode, nil)

	interactor := usecase.NewEpisodeInteractor(repository)
	episodeHandler := handlers.NewEpisodeHandler(*interactor)
	handler := http.HandlerFunc(episodeHandler.GetEpisodeByName)

	// Act
	var actual episodeDomain.Episode
	handler.ServeHTTP(rr, req)
	json.Unmarshal([]byte(rr.Body.String()), &actual)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, *episode, actual)
}

func TestEpisodeHandlerGetEpisodeByNameError(t *testing.T) {
	// Arrange
	t.Parallel()
	error := errors.New("Failure is apparently an option")
	episodeName := "Anasazi"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/episodes/"+episodeName, nil)
	if err != nil {
		t.Fatal(err)
	}

	repository := &repositoryMock.Episode{}
	repository.On(
		"GetEpisodeByName",
		episodeName,
	).Return(&episodeDomain.Episode{}, error)

	interactor := usecase.NewEpisodeInteractor(repository)
	episodeHandler := handlers.NewEpisodeHandler(*interactor)
	handler := http.HandlerFunc(episodeHandler.GetEpisodeByName)

	// Act
	handler.ServeHTTP(rr, req)
	actual, _ := ioutil.ReadAll(rr.Body)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Error retrieving episode details\n", string(actual))
}

func generateDomainEpisode(name string, season int, episode int) *episodeDomain.Episode {
	return &episodeDomain.Episode{
		Name:    name,
		Season:  season,
		Episode: episode,
	}
}
