package repository

import episodeDomain "github.com/napalm684/cleandemo/domain/episode"

// EpisodeRepository Interface
type EpisodeRepository interface {
	GetEpisodeByName(name string) (*episodeDomain.Episode, error)
}
