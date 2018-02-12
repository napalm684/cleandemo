package inmemory

import episodeDomain "github.com/napalm684/cleandemo/domain/episode"

//EpisodeRepository structure
type EpisodeRepository struct {
	episodes map[string]*episodeDomain.Episode
}

//NewEpisodeRepository returns a repository for episodes
func NewEpisodeRepository(episodes map[string]*episodeDomain.Episode) *EpisodeRepository {
	return &EpisodeRepository{
		episodes: episodes,
	}
}

//GetEpisodeByName returns an episode from in-memory data source
func (repo *EpisodeRepository) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	r := repo.episodes[name]
	return r, nil
}
