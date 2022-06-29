package deal

import (
	"euchre/deck"
	"testing"
)

func TestNewDeal(t *testing.T) {
	wantCards := 5
	wantKitty := 4

	d := deck.New()

	deal := New(d)

	if len(deal.hands[0]) != wantCards || len(deal.hands[1]) != wantCards || len(deal.hands[2]) != wantCards || len(deal.hands[3]) != wantCards {
		t.Fatal("Deal() dealt incorrect number of cards")
	}

	if len(deal.kitty) != wantKitty {
		t.Fatal("Deal() dealt incorrect number of cards in 'kitty'")
	}
}

func TestDealHand(t *testing.T) {
	want := 5

	d := deck.New()

	hands := New(d)

	hand := hands.Hand(0)

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
	d := deck.New()

	deal := New(d)

	kitty := deal.Kitty()

	if len(kitty) != want {
		t.Fatalf(`Kitty() returned %d. want: %d`, len(kitty), want)
	}
}
