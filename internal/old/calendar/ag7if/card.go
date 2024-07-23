package ag7if

import (
	cards "github.com/ag7if/playing-cards"
)

var weekDeck = cards.NewDeck()

func GetWeekCard(isoWeek int) cards.Card {
	return weekDeck[isoWeek-1]
}
