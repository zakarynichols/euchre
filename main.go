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

	// Give the players their cards.
	p := players.New(deal)

	deal.Pickup(p.PlayerOne, 0)

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

	winner, err := t.Winner()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Winner: ", winner)
}
