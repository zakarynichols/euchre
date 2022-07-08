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
			t.Fatalf(`NewDeck() dealt a malformed rank type: %d`, c.rank)
		}
		if totalRanks != wantRanks {
			t.Fatalf(`NewDeck() dealt %d of %d`, totalRanks, c.rank)
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

func TestNewDeal(t *testing.T) {
	wantCards := 5
	wantKitty := 4

	deal := New().Deal()

	if len(deal.hands[0]) != wantCards || len(deal.hands[1]) != wantCards || len(deal.hands[2]) != wantCards || len(deal.hands[3]) != wantCards {
		t.Fatal("Deal() dealt incorrect number of cards")
	}

	if len(deal.kitty) != wantKitty {
		t.Fatal("Deal() dealt incorrect number of cards in 'kitty'")
	}
}

func TestDealHand(t *testing.T) {
	want := 5

	d := New()

	hand := d.Deal().Hand(0)

	if len(hand) != want {
		t.Fatalf(`Hand(PlayerOne) has %d cards. want %d `, len(hand), want)
	}
}

// func TestDealHandError(t *testing.T) {
// 	d := deck.New()

// 	hands := New(d)

// 	_, err := hands.Hand("garb")

// 	if err != nil {
// 		if err.Error() != "invalid player key: garb" {
// 			t.Fatal("Hand('garb') returned invalid error message")
// 		}
// 	}
// }

func TestDealKitty(t *testing.T) {
	want := 4

	d := New()

	deal := d.Deal()

	kitty := deal.Kitty()

	if len(kitty) != want {
		t.Fatalf(`Kitty() returned %d. want: %d`, len(kitty), want)
	}
}
