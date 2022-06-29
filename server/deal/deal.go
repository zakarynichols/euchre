package deal

import (
	"euchre/deck"
)

// The deal package should only need the deck to deal the cards.
// Make sure to deal them sequentially and clock-wise just
// like a real euchre game.

type Deal struct {
	hands map[int][]deck.Card
	kitty []deck.Card
}

func New(d deck.Deck) Deal {
	return Deal{
		hands: map[int][]deck.Card{
			0: d[0:5],
			1: d[5:10],
			2: d[10:15],
			3: d[15:20],
		},
		kitty: d[20:24],
	}
}

func (d Deal) Hands() map[int][]deck.Card {
	return d.hands
}

func (d Deal) Hand(key int) []deck.Card {
	return d.hands[key]
}

func (d Deal) Kitty() []deck.Card {
	return d.kitty
}
