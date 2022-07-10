package players

import (
	"euchre/deck"
)

type Hand map[int]deck.Card

type Player struct {
	Key      PlayerKey
	IsDealer bool
	Placing  Placing
	hand     Hand
	Team     Team
}

type Team string

const (
	EmptyTeam Team = "Empty"
	RedTeam   Team = "Red"
	BlackTeam Team = "Black"
)

// Notice the pointer to a player. There are a couple methods that
// need to modify the receiver. In this case, the Player struct is
// the receiver. (p *Player) for example.
type Players map[PlayerKey]*Player

func (p *Player) SetDealer() {
	p.IsDealer = true
}

func (p *Player) SetTeam(t Team) {
	p.Team = t
}

func (p *Player) Swap(c *deck.Card, index int) *deck.Card {
	discard := p.hand[index]
	p.hand[index] = *c
	return &discard
}

// Type for all player strings.
// Prefer this type instead of naked strings.
// This doesn't prevent a consumer from passing
// a regular string. Just a helper when developing.
type PlayerKey int

const (
	// Players
	EmptyPlayer PlayerKey = 0
	One         PlayerKey = 1
	Two         PlayerKey = 2
	Three       PlayerKey = 3
	Four        PlayerKey = 4
)

type Placing int

const (
	None   Placing = 0
	First  Placing = 1
	Second Placing = 2
	Third  Placing = 3
	Fourth Placing = 4
)

func newPlayer(k PlayerKey, isDealer bool, p Placing, h Hand) *Player {
	return &Player{
		Key:      k,
		IsDealer: isDealer,
		Placing:  p,
		hand:     h,
	}
}

type Dealer interface {
	Hand(key int) []deck.Card
}

// Dealer returns a slice of cards. Convert the slice to a map
// to allow faster lookups, adds, and deletes, when playing the game.
func sliceToMap(d Dealer, index int) Hand {
	h := make(Hand, 0)

	hand := d.Hand(index)

	for i := range d.Hand(index) {
		h[i] = hand[i]
	}

	return h
}

func New(d Dealer) Players {
	return Players{
		One:   newPlayer(One, false, None, sliceToMap(d, 0)),
		Two:   newPlayer(Two, false, None, sliceToMap(d, 1)),
		Three: newPlayer(Three, false, None, sliceToMap(d, 2)),
		Four:  newPlayer(Four, false, None, sliceToMap(d, 3)),
	}
}

func (p Player) Hand() Hand {
	return p.hand
}

func (h Hand) Len() int {
	return len(h)
}
