package handler

import (
	"github.com/alinurmyrzakhanov/exercises/exercise4/judge/internal/ticTacToe/game"
)

type Handler struct {
	game *game.Game
}

func New() *Handler {
	return &Handler{
		game: game.New(),
	}
}
