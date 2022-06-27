package deck

import (
	"reflect"
	"testing"
)

var deck = NewDeck()

func TestDeckCards(t *testing.T) {
	want := 24

	if len(deck) != want {
		t.Fatalf(`NewDeck() = %d. want new deck length to equal %d.`, len(deck), want)
	}
}

func TestDeckContents(t *testing.T) {
	wantSuits := 4

	wantRanks := 6

	cards := make(map[int]Card)

	var totalSuits int
	var totalRanks int

	for n := 0; n < len(suits); n++ {
		totalSuits++
		for i := 0; i < len(ranks); i++ {
			cards[n] = Card{Suit: suits[n], Rank: ranks[n]}
			totalRanks++
		}
	}

	totalRanks = totalRanks / totalSuits

	for _, r := range cards {
		if totalSuits != wantSuits {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalSuits, r.Suit)
		}
		if r.Rank != NINE && r.Rank != TEN && r.Rank != JACK && r.Rank != QUEEN && r.Rank != KING && r.Rank != ACE {
			t.Fatalf(`NewDeck() dealt %s rank`, r.Rank)
		}
		if totalRanks != wantRanks {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalRanks, r.Rank)
		}
	}
}

func TestDeckShuffle(t *testing.T) {
	want := false

	if reflect.DeepEqual(deck, deck.Shuffle()) != want {
		t.Fatal("Shuffle() returned a deep equal deck")
	}
}
