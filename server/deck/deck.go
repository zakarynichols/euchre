package deck

import (
	"fmt"
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

	// Players
	PlayerOne   Player = "PlayerOne"
	PlayerTwo   Player = "PlayerTwo"
	PlayerThree Player = "PlayerThree"
	PlayerFour  Player = "PlayerFour"

	// Cards leftover after deal are called the 'kitty'
	Kitty = "Kitty"
)

// Model of a playing card
type Card struct {
	rank string
	suit string
}

// Slice of cards
type Deck []Card

// Type for all player strings.
// Prefer this type instead of naked strings.
// This doesn't prevent a consumer from passing
// a regular string. Just a helper when developing.
type Player string

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

type Deal map[Player][]Card

func (d Deck) Deal() Deal {
	// Slice expressions allow concise card dealing
	return Deal{PlayerOne: d[0:5], PlayerTwo: d[5:10], PlayerThree: d[10:15], PlayerFour: d[15:20], Kitty: d[20:24]}
}

func (d Deal) Hand(key Player) ([]Card, error) {
	if key != PlayerOne && key != PlayerTwo && key != PlayerThree && key != PlayerFour {
		return nil, fmt.Errorf(`invalid player key: %s`, key)
	}

	return d[key], nil
}

func (d Deal) Kitty() []Card {
	return d[Kitty]
}
