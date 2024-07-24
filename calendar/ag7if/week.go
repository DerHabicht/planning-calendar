package ag7if

import (
	cards "github.com/ag7if/playing-cards"
	"github.com/pkg/errors"
)

var deck cards.Deck = cards.NewDeck()

func ComputeWeekPlayingCard(isoweek int) (cards.Card, error) {
	if isoweek < 1 || isoweek > 53 {
		return cards.Card{}, errors.Errorf("invaild week number: %d", isoweek)
	}

	return deck[isoweek+1], nil
}
