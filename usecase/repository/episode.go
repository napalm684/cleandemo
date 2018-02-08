package repository

import episodeDomain "github.com/napalm684/cleandemo/domain/episode"

// Episode Repository Interface
type Episode interface {
	GetEpisodeByName(name string) (*episodeDomain.Episode, error)
}
