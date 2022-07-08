package trick

import (
	"euchre/deck"
	"euchre/players"
	"testing"
)

func TestRightBowerWins(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	p.PlayerFour.SetLead()

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Jack, deck.Spade),
				Player: *p.PlayerOne,
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   deck.NewCard(deck.Nine, deck.Heart),
				Player: *p.PlayerThree,
			},
			3: {
				Card:   deck.NewCard(deck.Ace, deck.Diamond),
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Spade,
	}

	winner := tr.Winner()

	if winner.Card != deck.NewCard(deck.Jack, deck.Spade) {
		t.Fatal("The jack of spades should be the win the trick")
	}
}

func TestLeftBowerWins(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	p.PlayerFour.SetLead()

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: *p.PlayerOne,
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   deck.NewCard(deck.Nine, deck.Heart),
				Player: *p.PlayerThree,
			},
			3: {
				Card:   deck.NewCard(deck.Ace, deck.Diamond),
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Spade,
	}

	winner := tr.Winner()

	if winner.Card != deck.NewCard(deck.Jack, deck.Club) {
		t.Fatal("The jack of clubs should be the win the trick")
	}
}

func TestHighestTrumpNoBowers(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	p.PlayerFour.SetLead()

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Queen, deck.Spade),
				Player: *p.PlayerOne,
			},
			1: {
				Card:   deck.NewCard(deck.Ace, deck.Spade),
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: *p.PlayerThree,
			},
			3: {
				Card:   deck.NewCard(deck.King, deck.Diamond),
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Spade,
	}

	winner := tr.Winner()

	if winner.Card != deck.NewCard(deck.Ace, deck.Spade) {
		t.Fatal("The ace of spades should win the trick")
	}
}

func TestLeadPlayerTrumpWins(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	p.PlayerFour.SetLead()

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.King, deck.Heart),
				Player: *p.PlayerOne,
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: *p.PlayerThree,
			},
			3: {
				Card:   deck.NewCard(deck.King, deck.Diamond),
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Spade,
	}

	winner := tr.Winner()

	if winner.Card != deck.NewCard(deck.King, deck.Diamond) {
		t.Fatal("The king of diamonds should win the trick")
	}
}

func TestLeadPlayerTrumpGetsBeat(t *testing.T) {
	d := deck.New()

	p := players.New(d.Deal())

	p.PlayerFour.SetLead()

	tr := Trick{
		Cards: Play{
			0: {
				Card:   deck.NewCard(deck.Queen, deck.Spade),
				Player: *p.PlayerOne,
			},
			1: {
				Card:   deck.NewCard(deck.Jack, deck.Club),
				Player: *p.PlayerTwo,
			},
			2: {
				Card:   deck.NewCard(deck.Queen, deck.Diamond),
				Player: *p.PlayerThree,
			},
			3: {
				Card:   deck.NewCard(deck.King, deck.Spade),
				Player: *p.PlayerFour,
			},
		},
		Trump: deck.Spade,
	}

	winner := tr.Winner()

	if winner.Card != deck.NewCard(deck.King, deck.Spade) {
		t.Fatalf("Want card: %v Got card: %v", deck.NewCard(deck.King, deck.Spade), winner.Card)
	}
}
