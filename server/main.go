package main

import (
	"euchre/deal"
	"euchre/deck"
	"euchre/players"
	"log"
)

func main() {
	// Get a new deck.
	d := deck.New()

	// Shuffle the deck.
	d.Shuffle()

	// Deal cards into piles.
	hands := deal.New(d)

	// All the cards dealt into hands
	// ready to be handed to the players.
	for i, v := range hands.Hands() {
		log.Print("Player Index: ", i)
		log.Print("Player Value: ", v)
		log.Print("------------------")
	}

	// Dealing the cards also returns the kitty.
	log.Println("Kitty: ", hands.Kitty())

	// Give the players their cards.
	p := players.New(hands)

	// Set the dealer.
	p.SetDealer(players.Four)

	// Show player one p1
	p1 := p[players.One]

	// The hand is a map of cards
	log.Print("Hand before play: ", p1.ShowHand())

	// Play one of player ones cards
	p[players.One].Play(p1.ShowHand()[0])

	// Play uses `delete` under the hood to remove
	// a card from a players hand.
	log.Print("Hand after play: ", p1.ShowHand())
}
