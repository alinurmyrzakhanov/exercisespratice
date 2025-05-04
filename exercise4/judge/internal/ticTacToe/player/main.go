package player

import (
	"github.com/alinurmyrzakhanov/exercises/exercise4/judge/internal"
)

type Player struct {
	Name  string          `json:"name"`
	URL   string          `json:"url"`
	Token *internal.Token `json:"token,omitempty"`
}

func New(name string, url string) *Player {
	return &Player{
		Name:  name,
		URL:   url,
		Token: nil,
	}
}
