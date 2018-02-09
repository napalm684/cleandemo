package inmemory

import episodeDomain "github.com/napalm684/cleandemo/domain/episode"

//Episode structure
type Episode struct {
	episodes map[string]*episodeDomain.Episode
}

//NewEpisodeRepository returns a repository for episodes
func NewEpisodeRepository(episodes map[string]*episodeDomain.Episode) *Episode {
	return &Episode{
		episodes: episodes,
	}
}

//GetEpisodeByName returns an episode from in-memory data source
func (repo *Episode) GetEpisodeByName(name string) (*episodeDomain.Episode, error) {
	r := repo.episodes[name]
	return r, nil
}
