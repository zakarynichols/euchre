package deck

import (
	"math/rand"
	"time"
)

// Playing card ranks and suits
const (
	// Possibly implemented in future versions
	Seven = "Seven"
	Eight = "Eight"

	// Ranks
	Nine  = "Nine"
	Ten   = "Ten"
	Jack  = "Jack"
	Queen = "Queen"
	King  = "King"
	Ace   = "Ace"

	// Suits
	Heart   = "Heart"
	Diamond = "Diamond"
	Club    = "Club"
	Spade   = "Spade"
)

// Model of a playing card
type Card struct {
	Rank string
	Suit string
}

// Slice of cards
type Deck []Card

// Ranks as a fixed array
var ranks = [6]string{Nine, Ten, Jack, Queen, King, Ace}

// Suits as a fixed array
var suits = [4]string{Heart, Diamond, Club, Spade}

// New returns a new deck of cards.
// Ranks and suits are set internally.
func New() Deck {
	deck := make(Deck, 0)

	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Rank: ranks[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}

	return deck
}

// Shuffle the deck.
func (deck Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

type Deal struct {
	P1, P2, P3, P4, Kitty []Card
}

func (d Deck) Deal() Deal {
	return Deal{P1: d[0:5], P2: d[5:10], P3: d[10:15], P4: d[15:20], Kitty: d[20:24]}
}
