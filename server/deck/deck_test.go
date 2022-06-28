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
			cards[n] = Card{Suit: suits[n], Rank: ranks[n]}
			totalCards++
		}
	}

	totalRanks = totalCards / totalSuits

	for _, c := range cards {
		if c.Suit != Heart && c.Suit != Diamond && c.Suit != Spade && c.Suit != Club {
			t.Fatalf(`NewDeck() dealt a malformed suit type: %s`, c.Suit)
		}
		if totalSuits != wantSuits {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalSuits, c.Suit)
		}
		if c.Rank != Nine && c.Rank != Ten && c.Rank != Jack && c.Rank != Queen && c.Rank != King && c.Rank != Ace {
			t.Fatalf(`NewDeck() dealt a malformed rank type: %s`, c.Rank)
		}
		if totalRanks != wantRanks {
			t.Fatalf(`NewDeck() dealt %d of %s`, totalRanks, c.Rank)
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

func TestDeckDeal(t *testing.T) {
	wantCards := 5
	wantKitty := 4

	deck := New()

	deal := deck.Deal()

	if len(deal.P1) != wantCards || len(deal.P2) != wantCards || len(deal.P3) != wantCards || len(deal.P4) != wantCards {
		t.Fatal("Deal() dealt incorrect number of cards")
	}

	if len(deal.Kitty) != wantKitty {
		t.Fatal("Deal() dealt incorrect number of cards in 'kitty'")
	}
}
