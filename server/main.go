package main

import (
	"euchre/deck"
	"log"
)

func main() {
	// Get a new deck
	d := deck.New()

	// Shuffle the deck
	d.Shuffle()

	// Deal cards
	hands := d.Deal()

	// Peek at the kitty
	log.Println(hands.Kitty())

	// All the players current hands
	for i, v := range hands {
		log.Print("Player Index: ", i)
		log.Print("Player Value: ", v)
		log.Print("------------------")
	}
}
