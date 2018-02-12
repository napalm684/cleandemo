package usecase

import (
	"fmt"

	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	"github.com/napalm684/cleandemo/usecase/repository"
)

//EpisodeService provides consumer with the ability to work with episodes
type EpisodeService struct {
	repository repository.EpisodeRepository
}

//NewEpisodeService returns
func NewEpisodeService(episodeRepository repository.EpisodeRepository) *EpisodeService {
	return &EpisodeService{
		repository: episodeRepository,
	}
}

//GetEpisodeByName returns episode assigned to the given episode name
func (episodeInteractor *EpisodeService) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	episode, err := episodeInteractor.repository.GetEpisodeByName(name)

	if err != nil {
		return nil, fmt.Errorf("Unable to fetch episode data: %s", err)
	}

	return episode, nil
}
