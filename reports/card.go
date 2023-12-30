package reports

import (
	"fmt"
	"strconv"

	"github.com/derhabicht/planning-calendar/calendar"
)

func MarshallCard(card calendar.Card) string {
	var suit string

	switch card.Suit {
	case calendar.Hearts:
		suit = `$\heartsuit$`
	case calendar.Clubs:
		suit = `$\clubsuit$`
	case calendar.Diamonds:
		suit = `$\diamondsuit$`
	case calendar.Spades:
		suit = `$\spadesuit$`
	default:
		suit = ""
	}

	var rank string
	switch card.Rank {
	case 0:
		rank = "JK"
	case 1:
		rank = "A"
	case 11:
		rank = "J"
	case 12:
		rank = "Q"
	case 13:
		rank = "K"
	default:
		rank = strconv.Itoa(int(card.Rank))
	}

	return fmt.Sprintf("%s%s", rank, suit)
}
