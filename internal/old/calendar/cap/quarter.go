package cap

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar/ag7if"
)

type Q int

const (
	Q1 Q = iota + 1
	Q2
	Q3
	Q4
)

func (q Q) String() string {
	switch q {
	case Q1:
		return "Q1"
	case Q2:
		return "Q2"
	case Q3:
		return "Q3"
	case Q4:
		return "Q4"
	default:
		panic(fmt.Errorf("%d is not a valid quarter", q))
	}
}

type Quarter struct {
	Quarter    Q
	FiscalYear int
}

func NewQuarter(d date.Date) Quarter {
	switch d.Month() {
	case time.October, time.November, time.December:
		return Quarter{Quarter: Q1, FiscalYear: d.Year() + 1}
	case time.January, time.February, time.March:
		return Quarter{Quarter: Q2, FiscalYear: d.Year()}
	case time.April, time.May, time.June:
		return Quarter{Quarter: Q3, FiscalYear: d.Year()}
	case time.July, time.August, time.September:
		return Quarter{Quarter: Q4, FiscalYear: d.Year()}
	default:
		panic(fmt.Errorf("%d is not a valid month", d.Month()))
	}
}

func (q Quarter) CyQuarter() ag7if.Quarter {
	switch q.Quarter {
	case Q1:
		return ag7if.NewQuarter(ag7if.Q4, q.FiscalYear-1)
	case Q2:
		return ag7if.NewQuarter(ag7if.Q1, q.FiscalYear)
	case Q3:
		return ag7if.NewQuarter(ag7if.Q2, q.FiscalYear)
	case Q4:
		return ag7if.NewQuarter(ag7if.Q3, q.FiscalYear)
	default:
		panic(fmt.Errorf("%d is not a valid quarter", q))
	}
}

func (q Quarter) NextQuarter() Quarter {
	switch q.Quarter {
	case Q1:
		return Quarter{Q2, q.FiscalYear}
	case Q2:
		return Quarter{Q3, q.FiscalYear}
	case Q3:
		return Quarter{Q4, q.FiscalYear}
	case Q4:
		return Quarter{Q1, q.FiscalYear + 1}
	default:
		panic(fmt.Errorf("%d is not a valid quarter", q))
	}

}

func (q Quarter) StartDate() date.Date {
	switch q.Quarter {
	case Q1:
		return date.New(q.FiscalYear-1, time.October, 1)
	case Q2:
		return date.New(q.FiscalYear, time.January, 1)
	case Q3:
		return date.New(q.FiscalYear, time.April, 1)
	case Q4:
		return date.New(q.FiscalYear, time.July, 1)
	default:
		panic(fmt.Errorf("%d is not a valid quarter", q))
	}
}

func (q Quarter) String() string {
	return fmt.Sprintf("FY%d%s", q.FiscalYear, q.Quarter)
}
