package deck

import (
	"math/rand"
	"time"
)

type Rank int
type Suit string

// Playing card ranks and suits
const (
	// Possibly implemented in future versions
	Seven = "Seven"
	Eight = "Eight"

	// Ranks
	Nine  Rank = 0
	Ten   Rank = 1
	Jack  Rank = 2
	Queen Rank = 3
	King  Rank = 4
	Ace   Rank = 5

	// Suits
	Heart   Suit = "Heart"
	Diamond Suit = "Diamond"
	Club    Suit = "Club"
	Spade   Suit = "Spade"

	// Cards leftover after deal are called the 'kitty'
	Kitty = "Kitty"
)

// Model of a playing card
type Card struct {
	rank Rank
	suit Suit
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) Suit() Suit {
	return c.suit
}

// Slice of cards
type Deck []Card

// Ranks as a fixed array
var ranks = [6]Rank{Nine, Ten, Jack, Queen, King, Ace}

// Suits as a fixed array
var suits = [4]Suit{Heart, Diamond, Club, Spade}

// New returns a new deck of cards.
// Ranks and suits are set internally.
func New() Deck {
	deck := make(Deck, 0)

	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				rank: ranks[i],
				suit: suits[n],
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
	hands map[int][]Card
	kitty []Card
}

// Deal the cards and kitty. The hands are a map of slices. The kitty is a slice.
func (d Deck) Deal() Deal {
	return Deal{
		hands: map[int][]Card{
			0: d[0:5],
			1: d[5:10],
			2: d[10:15],
			3: d[15:20],
		},
		kitty: d[20:24],
	}
}

// Shows all dealt hands.
func (d Deal) Hands() map[int][]Card {
	return d.hands
}

// Shows a dealt hand by a key.
func (d Deal) Hand(key int) []Card {
	return d.hands[key]
}

// Returns the cards in the kitty.
func (d Deal) Kitty() []Card {
	return d.kitty
}

type Swapper interface {
	Swap(c *Card, index int) *Card
}

// Called when the flipped over card in the kitty is picked up by the dealer.
func (d Deal) Pickup(s Swapper, index int) {
	k := d.kitty[0]
	discard := s.Swap(&k, index)
	d.kitty[0] = *discard
}
