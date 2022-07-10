package trick

import (
	"euchre/deck"
	"euchre/players"
	"testing"
)

func TestRightBowerWins(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Jack, deck.Spade),
				Player: p[players.One],
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.Nine, deck.Heart),
				Player: p[players.Three],
			},
			3: {
				Card:   deck.NewCard(deck.Ace, deck.Diamond),
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	tr.SetLead(p[players.Four])

	winner, _ := tr.Winner()

	if winner.Card != deck.NewCard(deck.Jack, deck.Spade) {
		t.Fatal("The jack of spades should be the win the trick")
	}
}

func TestLeftBowerWins(t *testing.T) {
	want := deck.NewCard(deck.Jack, deck.Club)

	d := deck.New()

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: p[players.One],
			},
			1: {
				Card:   want,
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.Nine, deck.Heart),
				Player: p[players.Three],
			},
			3: {
				Card:   deck.NewCard(deck.Ace, deck.Diamond),
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	tr.SetLead(p[players.Four])

	got, _ := tr.Winner()

	if got.Card != want {
		t.Fatalf("Got %v; Want %v", got.Card, want)
	}
}

func TestHighestTrumpNoBowers(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Queen, deck.Spade),
				Player: p[players.One],
			},
			1: {
				Card:   deck.NewCard(deck.Ace, deck.Spade),
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: p[players.Three],
			},
			3: {
				Card:   deck.NewCard(deck.Ace, deck.Diamond),
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	tr.SetLead(p[players.Four])

	winner, _ := tr.Winner()

	if winner.Card != deck.NewCard(deck.Ace, deck.Spade) {
		t.Fatal("The ace of spades should win the trick")
	}
}

func TestLeadDealerOffsuitWins(t *testing.T) {
	want := deck.NewCard(deck.King, deck.Diamond)

	d := deck.New()

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.King, deck.Heart),
				Player: p[players.One],
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Diamond),
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: p[players.Three],
			},
			3: {
				Card:   want,
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	tr.SetLead(p[players.Four])

	got, _ := tr.Winner()

	if got.Card != want {
		t.Fatalf("Got %v; Want %v", got.Card, want)
	}
}

func TestLeadPlayerTrumpGetsBeat(t *testing.T) {
	d := deck.New()

	want := deck.NewCard(deck.Jack, deck.Club)

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Queen, deck.Spade),
				Player: p[players.One],
			},
			1: {
				Card:   want,
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: p[players.Three],
			},
			3: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	tr.SetLead(p[players.Four])

	got, _ := tr.Winner()

	if got.Card != want {
		t.Fatalf("Got card: %v Want card: %v", got.Card, want)
	}
}

func TestErrNoLeaderDealer(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Queen, deck.Spade),
				Player: p[players.One],
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: p[players.Two],
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: p[players.Three],
			},
			3: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: p[players.Four],
			},
		},
		Trump: deck.Spade,
	}

	_, err := tr.Winner()

	if err != ErrNoLeadDealer {
		t.Fatalf("got error: %v. want error: %v", err, ErrNoLeadDealer)
	}
}
