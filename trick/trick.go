package trick

import (
	"errors"
	"euchre/deck"
	"euchre/players"
)

var (
	ErrNoLeadDealer = errors.New("a lead dealer has not been set")
	ErrNoWinner     = errors.New("could not determine a winner")
)

var (
	EmptyCard   = deck.NewCard(deck.EmptyRank, deck.EmptySuit)
	EmptyPlayer = players.Player{}
)

type Trick struct {
	Cards Play
	Trump deck.Suit
	Lead  *players.Player
}

type Play map[int]struct {
	Card   deck.Card
	Player *players.Player
}

type Winner struct {
	Card   deck.Card
	Player *players.Player
}

func (t *Trick) SetLead(p *players.Player) {
	t.Lead = p
}

func (t Trick) isLeadSet() error {
	if t.Lead == nil {
		return ErrNoLeadDealer
	}
	return nil
}

func (t Trick) leftBowerSuit() deck.Suit {
	if t.Trump == deck.Club {
		return deck.Spade
	}
	if t.Trump == deck.Spade {
		return deck.Club
	}
	if t.Trump == deck.Diamond {
		return deck.Heart
	}
	if t.Trump == deck.Heart {
		return deck.Diamond
	}
	return deck.EmptySuit
}

// The new trump suit determined by the lead player of the trick.
// Only gets set if there are no trump cards in the trick.
func (t Trick) newTrump(hasTrump bool) deck.Suit {
	var newTrump deck.Suit
	// If there aren't any trump cards; the new trump is set to the lead players suit
	for _, v := range t.Cards {
		if (t.Lead == v.Player) && !hasTrump {
			newTrump = v.Card.Suit
		}
	}

	return newTrump
}

func (t Trick) hasTrump(leftBowerSuit deck.Suit) bool {
	// If there are any trump cards in play this gets set to true
	hasTrump := false

	// Determine if there are any trump cards in the trick
	for _, v := range t.Cards {
		if v.Card.Suit == t.Trump ||
			// The left bower is considered a trump card
			(v.Card.Suit == leftBowerSuit &&
				v.Card.Rank == deck.Jack) {
			hasTrump = true
		}
	}

	return hasTrump
}

func (trick Trick) Winner() (Winner, error) {
	err := trick.isLeadSet()

	if err != nil {
		return Winner{
			Card:   EmptyCard,
			Player: &players.Player{},
		}, err
	}

	leftBowerSuit := trick.leftBowerSuit()

	hasTrump := trick.hasTrump(leftBowerSuit)

	newTrump := trick.newTrump(hasTrump)

	winner := Winner{
		Card:   EmptyCard,
		Player: &players.Player{},
	}

	for _, v := range trick.Cards {
		// If the trick contains trump cards
		if hasTrump {
			// Right bower
			if v.Card.Rank == deck.Jack &&
				v.Card.Suit == trick.Trump {
				winner = Winner{
					v.Card,
					v.Player,
				}
				// The right bower is the highest ranked card in a trick.
				// Break here since there is no need to iterate further.
				break
			}

			// Left bower
			if v.Card.Suit != trick.Trump &&
				v.Card.Suit == leftBowerSuit &&
				v.Card.Rank == deck.Jack {
				winner = Winner{
					v.Card,
					v.Player,
				}
			}

			// Highest trump card
			if v.Card.Suit == trick.Trump {
				// If we find a trump card; set it only if the winner is still a zero-value.
				if winner.Card == EmptyCard {
					winner = Winner{
						v.Card,
						v.Player,
					}
				}
				// We only want to set 'winner' again if the current card rank is higher.
				if v.Card.Rank > winner.Card.Rank {
					if winner.Card.Rank == deck.Jack {
						// Break if the winnning card is a jack.
						break
					}
					winner = Winner{
						v.Card,
						v.Player,
					}
				}
			}
			// Fallback to using lead player suit for trump
		} else {
			if v.Card.Suit == newTrump {
				// If we find a trump card; set it only if the winner is still a zero-value.
				if winner.Card == EmptyCard {
					winner = Winner{
						v.Card,
						v.Player,
					}
				}
				if v.Card.Rank > winner.Card.Rank {
					// We only want to set 'winner' again if the current card rank is higher.
					winner = Winner{
						v.Card,
						v.Player,
					}
				}
			}
		}
	}

	if winner.Card == EmptyCard || winner.Player.Key == EmptyPlayer.Key {
		return Winner{EmptyCard, &players.Player{}}, ErrNoWinner
	}

	return winner, nil
}
