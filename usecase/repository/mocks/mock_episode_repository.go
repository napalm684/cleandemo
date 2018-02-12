package mocks

import (
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/stretchr/testify/mock"
)

//EpisodeRepository repository mock
type EpisodeRepository struct {
	mock.Mock
}

//GetEpisodeByName - Retrieves episode details using given name
func (episode *EpisodeRepository) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	args := episode.Called(name)

	return args.Get(0).(*episodeDomain.Episode), args.Error(1)
}
