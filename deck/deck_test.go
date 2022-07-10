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
			t.Fatalf(`NewDeck() dealt a malformed rank type: %d`, c.Rank)
		}
		if totalRanks != wantRanks {
			t.Fatalf(`NewDeck() dealt %d of %d`, totalRanks, c.Rank)
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

	if len(deal.Hands[0]) != wantCards || len(deal.Hands[1]) != wantCards || len(deal.Hands[2]) != wantCards || len(deal.Hands[3]) != wantCards {
		t.Fatal("Deal() dealt incorrect number of cards")
	}

	if len(deal.Kitty) != wantKitty {
		t.Fatal("Deal() dealt incorrect number of cards in 'kitty'")
	}
}

func TestDealHand(t *testing.T) {
	want := 5

	d := New()

	hand := d.Deal().Hands[0]

	if len(hand) != want {
		t.Fatalf(`Hand(PlayerOne) has %d cards. want %d `, len(hand), want)
	}
}

func TestDealKitty(t *testing.T) {
	want := 4

	d := New()

	deal := d.Deal()

	kitty := deal.Kitty

	if len(kitty) != want {
		t.Fatalf(`Kitty() returned %d. want: %d`, len(kitty), want)
	}
}

// Mock swap and method to satisfy the Swapper interface
type Swap struct {
	Card
}

func (s *Swap) Swap(c *Card, i int) *Card {
	card := s.Card
	return &card
}
func TestPickup(t *testing.T) {
	want := NewCard(Jack, Diamond)

	// Get a new deck.
	d := New()

	// Deal cards into piles.
	deal := d.Deal()

	s := Swap{want}

	deal.Pickup(&s, 0)

	got := deal.Kitty[0]

	// Kitty should be the jack of diamonds after pickup
	if got != want {
		t.Fatalf(`Pickup() got %v. want %v`, got, want)
	}
}
