package main

import (
	"euchre/deck"
	"euchre/players"
	"euchre/trick"
	"fmt"
	"log"
)

func main() {
	// Get a new deck.
	newDeck := deck.New()

	// Shuffle the deck.
	newDeck.Shuffle()
	newDeck.Shuffle()
	newDeck.Shuffle()

	// Deal cards into piles.
	deal := newDeck.Deal()

	// Give the players their cards.
	p := players.New(deal)

	deal.Pickup(p[players.One], 0)

	// Set the dealer.
	p[players.One].SetDealer()

	p[players.One].SetTeam(players.BlackTeam)
	p[players.Two].SetTeam(players.BlackTeam)

	p[players.Three].SetTeam(players.RedTeam)
	p[players.Four].SetTeam(players.RedTeam)

	// Simulate 5 tricks
	var t trick.Trick
	var winners []trick.Winner
	var prevWinner trick.Winner
	for i := 0; i < 5; i++ {
		// An example trick
		t = trick.Trick{
			Cards: trick.Play{
				0: {
					Card:   p[players.One].Hand()[i],
					Player: p[players.One],
				},
				1: {
					Card:   p[players.Two].Hand()[i],
					Player: p[players.Two],
				},
				2: {
					Card:   p[players.Three].Hand()[i],
					Player: p[players.Three],
				},
				3: {
					Card:   p[players.Four].Hand()[i],
					Player: p[players.Four],
				},
			},
			Trump: deck.Spade,
		}

		if prevWinner.Player == nil {
			t.SetLead(p[players.One])
		} else {
			t.SetLead(prevWinner.Player)
		}

		winner, err := t.Winner()

		if err != nil {
			log.Fatal(err)
		}

		prevWinner = winner

		winners = append(winners, winner)
	}

	points := Points{}
	for _, v := range winners {
		fmt.Printf("Winning Player: %v \n", v.Player.Key)
		fmt.Printf("Winning card is %v of %v \n", v.Card.Rank, v.Card.Suit)
		fmt.Printf("Player is team %v \n", v.Player.Team)
		if v.Player.Team == players.RedTeam {
			points.Red++
		} else {
			points.Black++
		}
		fmt.Print("------------------------------------ \n")
	}

	var winningTeam string
	var winningTricks int
	if points.Black > points.Red {
		winningTeam = string(players.BlackTeam)
		winningTricks = points.Black
	} else {
		winningTeam = string(players.RedTeam)
		winningTricks = points.Red
	}

	fmt.Printf("%v team wins with %d tricks \n", winningTeam, winningTricks)
}

type Points struct {
	Red   int
	Black int
}
