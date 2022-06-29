package game

import (
	"euchre/players"
	"log"
)

// The game package holds the logic to actually play euchre.
// The dependencies should be fed into the game.
// Deps: deal and players

type Game struct {
	Players players.Players
}

func New(p players.Players, d players.PlayerKey) Game {

	return Game{Players: p}
}

func (g Game) Play() {

	log.Print("Play()...")
}
