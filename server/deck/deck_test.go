package deck

import "testing"

var deck = NewDeck()

func TestNewDeckTotalCards(t *testing.T) {
	want := 24

	if len(deck) != want {
		t.Fatalf(`NewDeck() = %d. want new deck length to equal %d.`, len(deck), want)
	}
}

func TestNewDeckSuit(t *testing.T) {
	want := 6

	type Suit struct {
		value   string
		howMany int
	}

	s := make(map[string]Suit)

	for n := 0; n < len(suits); n++ {
		for i := 0; i < len(ranks); i++ {
			s[suits[n]] = Suit{value: suits[n], howMany: i + 1}
		}
	}

	for _, suit := range s {
		if suit.howMany != want {
			t.Fatalf(`NewDeck() dealt %d of %s`, suit.howMany, suit.value)
		}
	}
}

func TestNewDeckRank(t *testing.T) {
	want := 4

	type Rank struct {
		value   string
		howMany int
	}

	r := make(map[string]Rank)

	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			r[ranks[i]] = Rank{value: ranks[i], howMany: n + 1}
		}
	}

	for _, rank := range r {
		if rank.howMany != want {
			t.Fatalf(`NewDeck() dealt %d of %s`, rank.howMany, rank.value)
		}

		if rank.value != "Nine" && rank.value != "Ten" && rank.value != "Jack" && rank.value != "Queen" && rank.value != "King" && rank.value != "Ace" {
			t.Fatalf(`NewDeck() dealt %s rank`, rank.value)
		}
	}
}
