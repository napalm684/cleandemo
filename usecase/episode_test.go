package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/napalm684/cleandemo/usecase"
	repositoryMock "github.com/napalm684/cleandemo/usecase/repository/mocks"
)

func TestGetEpisodeByName(t *testing.T) {
	// Arrange
	t.Parallel()
	episode := generateDomainEpisode("Kitten", 11, 6)

	repository := &repositoryMock.Episode{}
	repository.On(
		"GetEpisodeByName",
		episode.Name,
	).Return(episode, nil)

	interactor := usecase.NewEpisodeInteractor(repository)

	// Act
	result, err := interactor.GetEpisodeByName(episode.Name)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, episode, result)
	repository.AssertExpectations(t)
}

func TestGetEpisodeByNameError(t *testing.T) {
	// Arrange
	t.Parallel()
	error := errors.New("Unable to get episode by name")
	episodeName := "Anasazi"

	repository := &repositoryMock.Episode{}
	repository.On(
		"GetEpisodeByName",
		episodeName,
	).Return(&episodeDomain.Episode{}, error)

	interactor := usecase.NewEpisodeInteractor(repository)

	// Act
	result, err := interactor.GetEpisodeByName(episodeName)

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, "Unable to fetch episode data: "+error.Error(), err.Error())
	repository.AssertExpectations(t)
}

func generateDomainEpisode(name string, season int, episode int) *episodeDomain.Episode {
	return &episodeDomain.Episode{
		Name:    name,
		Season:  season,
		Episode: episode,
	}
}
