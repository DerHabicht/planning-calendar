package ag7if

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

type Q int

const (
	Q1 Q = iota + 1
	Q2
	Q3
	Q4
)

func (q Q) StartWeek() int {
	switch q {
	case Q1:
		return 1
	case Q2:
		return 14
	case Q3:
		return 27
	case Q4:
		return 40
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF quarter", q))
	}
}

func (q Q) StartMonth() time.Month {
	switch q {
	case Q1:
		return time.January
	case Q2:
		return time.April
	case Q3:
		return time.June
	case Q4:
		return time.October
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF quarter", q))
	}
}

func NewQ(isoWeek int) Q {
	switch {
	case (1 <= isoWeek) && (isoWeek <= 13):
		return Q1
	case (14 <= isoWeek) && (isoWeek <= 26):
		return Q2
	case (27 <= isoWeek) && (isoWeek <= 39):
		return Q3
	case (40 <= isoWeek) && (isoWeek <= 53):
		return Q4
	default:
		panic(fmt.Errorf("%d is not a valid week number", isoWeek))
	}
}

type Quarter struct {
	Quarter Q
	Year    int
}

func NewQuarter(date date.Date) Quarter {
	year, week := date.ISOWeek()
	qtr := NewQ(week)

	return Quarter{
		Quarter: qtr,
		Year:    year,
	}
}

func (q Quarter) NextQuarter() Quarter {
	switch q.Quarter {
	case Q1:
		return Quarter{Quarter: Q2, Year: q.Year}
	case Q2:
		return Quarter{Quarter: Q3, Year: q.Year}
	case Q3:
		return Quarter{Quarter: Q4, Year: q.Year}
	case Q4:
		return Quarter{Quarter: Q1, Year: q.Year + 1}
	default:
		panic(errors.Errorf("invalid Quarter value: %d", q))
	}
}

func (q Quarter) String() string {
	return fmt.Sprintf("CY%d%s", q.Year, q.Quarter)
}
