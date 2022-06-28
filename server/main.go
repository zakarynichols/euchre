package main

import (
	"euchre/deck"
)

func main() {
	deck := deck.New()
	deck.Shuffle()
	deck.Deal()

}
