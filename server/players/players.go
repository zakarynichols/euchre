package players

import (
	"euchre/deck"
)

type Hand map[int]deck.Card

type Player struct {
	key      PlayerKey
	isDealer bool
	placing  Placing
	hand     Hand
}

// Notice the pointer to a player
type Players map[PlayerKey]*Player

func (p Players) SetDealer(key PlayerKey) {
	// Reassignment is possible thanks to being a pointer.
	// Will not compile if player is a raw struct value.
	p[key].isDealer = true
}

// Type for all player strings.
// Prefer this type instead of naked strings.
// This doesn't prevent a consumer from passing
// a regular string. Just a helper when developing.
type PlayerKey string

const (
	// Players
	One   PlayerKey = "PlayerOne"
	Two   PlayerKey = "PlayerTwo"
	Three PlayerKey = "PlayerThree"
	Four  PlayerKey = "PlayerFour"
)

type Placing int

const (
	None   Placing = 0
	First  Placing = 1
	Second Placing = 2
	Third  Placing = 3
	Fourth Placing = 4
)

func NewPlayer(k PlayerKey, isDealer bool, p Placing, h Hand) *Player {
	return &Player{
		key:      k,
		isDealer: isDealer,
		placing:  p,
		hand:     h,
	}
}

// Dealer 'hands' the cards to a player.
type Dealer interface {
	Hand(key int) []deck.Card
}

// Dealer returns a slice of cards. Convert the slice to a map
// to allow faster lookups, adds, and deletes, when playing the game.
func NewHand(d Dealer, index int) Hand {
	m := make(map[int]deck.Card, 0)

	h := d.Hand(index)

	for i := range d.Hand(index) {
		m[i] = h[i]
	}

	return m
}

func New(d Dealer) Players {
	return Players{
		One:   NewPlayer(One, false, None, NewHand(d, 0)),
		Two:   NewPlayer(Two, false, None, NewHand(d, 1)),
		Three: NewPlayer(Three, false, None, NewHand(d, 2)),
		Four:  NewPlayer(Four, false, None, NewHand(d, 3)),
	}
}

func (p Player) ShowHand() Hand {
	return p.hand
}

func (p Player) Play(c deck.Card) {
	for i, v := range p.hand {
		if c == v {
			delete(p.hand, i)
		}
	}
}
