package player

import (
	"github.com/alinurmyrzakhanov/exerciseso/exercise4/judge/internal"
)

func (p *Player) SetToken(token internal.Token) *Player {
	cp := *p
	cp.Token = &token
	return &cp
}
