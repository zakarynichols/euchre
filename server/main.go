package main

import (
	"euchre/deck"
	"euchre/players"
	"euchre/trick"
	"log"
)

func main() {
	// Get a new deck.
	newDeck := deck.New()

	// Shuffle the deck.
	newDeck.Shuffle()

	// Deal cards into piles.
	deal := newDeck.Deal()

	// All the cards dealt into hands
	// ready to be handed to the players.
	// for i, v := range deal.Hands() {
	// 	log.Print("Player Index: ", i)
	// 	log.Print("Player Value: ", v)
	// 	log.Print("------------------")
	// }

	// Give the players their cards.
	p := players.New(deal)

	log.Print(deal.Kitty())
	log.Print(p.PlayerOne.Hand())

	deal.Pickup(p.PlayerOne, 0)

	log.Print(deal.Kitty())
	log.Print(p.PlayerOne.Hand())

	// Set the dealer.
	p.PlayerOne.SetDealer()

	// Set lead player of the trick
	p.PlayerTwo.SetLead()

	// An example trick
	t := trick.Trick{
		Cards: trick.Play{
			0: {
				Card:   p.PlayerOne.Hand()[0],
				Player: *p.PlayerOne,
			},
			1: {
				Card:   p.PlayerTwo.Hand()[0],
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   p.PlayerThree.Hand()[0],
				Player: *p.PlayerThree,
			},
			3: {
				Card:   p.PlayerFour.Hand()[0],
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Heart,
	}

	winner := t.Winner()

	log.Print(p.PlayerTwo.Lead())

	log.Print("Winner: ", winner)
}
