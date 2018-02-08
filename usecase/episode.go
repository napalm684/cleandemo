package usecase

import (
	"fmt"

	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/napalm684/cleandemo/usecase/repository"
)

//EpisodeInteractor provides consumer with the ability to work with episodes
type EpisodeInteractor struct {
	repository repository.Episode
}

//NewEpisodeInteractor returns
func NewEpisodeInteractor(episodeRepository repository.Episode) *EpisodeInteractor {
	return &EpisodeInteractor{
		repository: episodeRepository,
	}
}

//GetEpisodeByName returns episode assigned to the given episode name
func (episodeInteractor *EpisodeInteractor) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	episode, err := episodeInteractor.repository.GetEpisodeByName(name)

	if err != nil {
		return nil, fmt.Errorf("Unable to fetch episode data: %s", err)
	}

	return episode, nil
}
