package players

import (
	"euchre/deck"
	"fmt"
)

type Hand map[int]deck.Card

type Player struct {
	key      PlayerKey
	isDealer bool
	placing  Placing
	hand     Hand
	isLead   bool
}

// Notice the pointer to a player. There are a couple methods that
// need to modify the receiver. In this case, the Player struct is
// the receiver. (p *Player) for example.
type Players struct {
	PlayerOne   *Player
	PlayerTwo   *Player
	PlayerThree *Player
	PlayerFour  *Player
}

func (p *Player) SetDealer() {
	p.isDealer = true
}

func (p *Player) SetLead() {
	p.isLead = true
}

func (p *Player) Swap(c *deck.Card, index int) *deck.Card {
	discard := p.hand[index]
	p.hand[index] = *c
	return &discard
}

func (p Player) Lead() bool {
	return p.isLead
}

// Type for all player strings.
// Prefer this type instead of naked strings.
// This doesn't prevent a consumer from passing
// a regular string. Just a helper when developing.
type PlayerKey string

const (
	// Players
	EmptyPlayer PlayerKey = "EmptyPlayer"
	One         PlayerKey = "PlayerOne"
	Two         PlayerKey = "PlayerTwo"
	Three       PlayerKey = "PlayerThree"
	Four        PlayerKey = "PlayerFour"
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
		PlayerOne:   newPlayer(One, false, None, sliceToMap(d, 0)),
		PlayerTwo:   newPlayer(Two, false, None, sliceToMap(d, 1)),
		PlayerThree: newPlayer(Three, false, None, sliceToMap(d, 2)),
		PlayerFour:  newPlayer(Four, false, None, sliceToMap(d, 3)),
	}
}

func (p Player) Hand() Hand {
	return p.hand
}

func (p Player) Play(c deck.Card) error {
	found := false

	for i, v := range p.hand {
		if c == v {
			found = true
			delete(p.hand, i)
			return nil
		}
	}

	if !found {
		return fmt.Errorf("Player %s does not have card %v", p.key, c)
	}

	return nil
}

func (h Hand) Len() int {
	return len(h)
}

func (p Player) Key() PlayerKey {
	return p.key
}
