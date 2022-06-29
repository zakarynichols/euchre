package deck

import (
	"reflect"
	"testing"
)

func TestDeckTotalCards(t *testing.T) {
	deck := New()

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
	var totalCards int
	var totalRanks int

	for n := 0; n < len(suits); n++ {
		totalSuits++
		for i := 0; i < len(ranks); i++ {
			cards[n] = Card{suit: suits[n], rank: ranks[n]}
			totalCards++
		}
	}

	totalRanks = totalCards / totalSuits

	for _, c := range cards {
		if c.suit != Heart && c.suit != Diamond && c.suit != Spade && c.suit != Club {
			t.Fatalf(`NewDeck() dealt a malformed suit type: %s`, c.suit)
		}
		if totalSuits != wantSuits {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalSuits, c.suit)
		}
		if c.rank != Nine && c.rank != Ten && c.rank != Jack && c.rank != Queen && c.rank != King && c.rank != Ace {
			t.Fatalf(`NewDeck() dealt a malformed rank type: %s`, c.rank)
		}
		if totalRanks != wantRanks {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalRanks, c.rank)
		}
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := New()

	newDeck := New()

	if reflect.DeepEqual(deck, newDeck) != true {
		t.Fatal("New decks are not equal")
	}

	if reflect.DeepEqual(deck.Shuffle(), newDeck) != false {
		t.Fatal("New deck failed to shuffle")
	}
}
