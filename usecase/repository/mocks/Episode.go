package mocks

import (
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/stretchr/testify/mock"
)

//Episode repository mock
type Episode struct {
	mock.Mock
}

//GetEpisodeByName - Retrieves episode details using given name
func (episode *Episode) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	episode.Called(name)

	return &episodeDomain.Episode{
		Name:    name,
		Season:  11,
		Episode: 6,
	}, nil
}
