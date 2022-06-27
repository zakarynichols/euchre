package deck

import (
	"math/rand"
	"time"
)

// These will not be implemented in the original service.
const SEVEN = "Seven"
const EIGHT = "Eight"

// Ranks to play euchre.
const NINE = "Nine"
const TEN = "Ten"
const JACK = "Jack"
const QUEEN = "Queen"
const KING = "King"
const ACE = "Ace"

// Suits
const HEART = "Heart"
const DIAMOND = "Diamond"
const CLUB = "Club"
const SPADE = "Spade"

// Model of a playing card.
type Card struct {
	Rank string
	Suit string
}

type Deck []Card

func (d Deck) Shuffle() Deck {
	deck := make(Deck, 0)

	deck = append(deck, d...)

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

var ranks = [6]string{NINE, TEN, JACK, QUEEN, KING, ACE}
var suits = [4]string{HEART, DIAMOND, CLUB, SPADE}

// NewDeck returns a new deck of cards.
// Ranks and suits are set internally.
func NewDeck() Deck {
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
