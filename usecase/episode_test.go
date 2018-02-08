package usecase_test

import (
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

func generateDomainEpisode(name string, season int, episode int) *episodeDomain.Episode {
	return &episodeDomain.Episode{
		Name:    name,
		Season:  season,
		Episode: episode,
	}
}
