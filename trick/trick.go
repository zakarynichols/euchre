package trick

import (
	"errors"
	"euchre/deck"
	"euchre/players"
)

var (
	ErrNoLeadDealer = errors.New("a lead dealer has not been set")
	ErrNoWinner     = errors.New("could dot determine a winner")
)

var (
	EmptyCard = deck.NewCard(deck.EmptyRank, deck.EmptySuit)
)

type Trick struct {
	Cards Play
	Trump deck.Suit
}

type Play map[int]struct {
	Card   deck.Card
	Player players.Player
}

type Winner struct {
	Card   deck.Card
	Player players.PlayerKey
}

func (t Trick) isLeadSet() error {
	isLeadSet := false
	for _, v := range t.Cards {
		if v.Player.Lead() {
			isLeadSet = true
		}
	}
	if !isLeadSet {
		return ErrNoLeadDealer
	}
	return nil
}

func (t Trick) leftBowerSuit() deck.Suit {
	var leftBowerSuit deck.Suit

	if t.Trump == deck.Club {
		leftBowerSuit = deck.Spade
	}
	if t.Trump == deck.Spade {
		leftBowerSuit = deck.Club
	}
	if t.Trump == deck.Diamond {
		leftBowerSuit = deck.Heart
	}
	if t.Trump == deck.Heart {
		leftBowerSuit = deck.Diamond
	}

	return leftBowerSuit
}

// The new trump suit determined by the lead player of the trick.
// Only gets set if there are no trump cards in the trick.
func (t Trick) newTrump(hasTrump bool) deck.Suit {
	var newTrump deck.Suit
	// If there aren't any trump cards; the new trump is set to the lead players suit
	for _, v := range t.Cards {
		if v.Player.Lead() && !hasTrump {
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
			Player: players.EmptyPlayer,
		}, err
	}

	leftBowerSuit := trick.leftBowerSuit()

	hasTrump := trick.hasTrump(leftBowerSuit)

	newTrump := trick.newTrump(hasTrump)

	winner := Winner{
		Card:   EmptyCard,
		Player: players.EmptyPlayer,
	}

	for _, v := range trick.Cards {
		// If the trick contains trump cards
		if hasTrump {
			// Right bower
			if v.Card.Rank == deck.Jack &&
				v.Card.Suit == trick.Trump {
				winner = Winner{
					v.Card,
					v.Player.Key(),
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
					v.Player.Key(),
				}
			}

			// Highest trump card
			if v.Card.Suit == trick.Trump {
				// If we find a trump card; set it only if the winner is still a zero-value.
				if winner.Card == EmptyCard {
					winner = Winner{
						v.Card,
						v.Player.Key(),
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
						v.Player.Key(),
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
						v.Player.Key(),
					}
				}
				if v.Card.Rank > winner.Card.Rank {
					// We only want to set 'winner' again if the current card rank is higher.
					winner = Winner{
						v.Card,
						v.Player.Key(),
					}
				}
			}
		}
	}

	if winner.Card == EmptyCard || winner.Player == players.EmptyPlayer {
		return Winner{EmptyCard, players.EmptyPlayer}, ErrNoWinner
	}

	return winner, nil
}
