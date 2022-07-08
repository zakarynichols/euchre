package trick

import (
	"euchre/deck"
	"euchre/players"
	"log"
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

func (t Trick) Winner() Winner {
	// Don't determine a winner unless a leader has been set.
	// Will be moved into its own function.
	var isLeadSet bool
	for _, v := range t.Cards {
		if v.Player.Lead() {
			isLeadSet = true
		}
	}
	if !isLeadSet {
		log.Fatal("A lead dealer has not been set.")
	}

	var winner Winner

	// The new trump suit determined by the lead player of the trick.
	// Only gets set if there are no trump cards in the trick.
	var newTrump deck.Suit

	var leftBowerSuit deck.Suit

	// If there are any trump cards in play this gets set to true
	var hasTrump bool

	// Is there a pattern that expresses this logic better?
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

	// Determine if there are any trump cards in the trick
	for _, v := range t.Cards {
		if v.Card.Suit == t.Trump {
			hasTrump = true
		}
	}

	// If there aren't any trump cards; the new trump is set to the lead players suit
	for _, v := range t.Cards {
		if v.Player.Lead() && !hasTrump {
			newTrump = v.Card.Suit
		}
	}

	for _, v := range t.Cards {
		// If the trick contains trump cards
		if hasTrump {

			// Right bower
			if v.Card.Rank == deck.Jack &&
				v.Card.Suit == t.Trump {
				winner = Winner{
					v.Card,
					v.Player.Key(),
				}
				// The right bower is the highest ranked card in a trick.
				// Break here since there is no need to iterate further.
				break
			}

			// Left bower
			if v.Card.Suit != t.Trump &&
				v.Card.Suit == leftBowerSuit &&
				v.Card.Rank == deck.Jack {
				winner = Winner{
					v.Card,
					v.Player.Key(),
				}
			}

			// Highest trump card
			if v.Card.Suit == t.Trump {
				// If we find a trump card; set it only if the winner is still a zero-value.
				if winner.Card == *new(deck.Card) {
					winner = Winner{
						v.Card,
						v.Player.Key(),
					}
				}
				// We only want to set 'winner' again if the current card rank is higher.
				if v.Card.Rank > winner.Card.Rank {
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
				if winner.Card == *new(deck.Card) {
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

	return winner
}
