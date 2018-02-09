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
	args := episode.Called(name)

	return args.Get(0).(*episodeDomain.Episode), args.Error(1)
}
