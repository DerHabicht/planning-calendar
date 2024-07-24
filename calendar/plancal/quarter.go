package plancal

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/ag7if"
	"github.com/derhabicht/planning-calendar/calendar/cap"
)

type QT int

const (
	CalendarQuarter QT = iota
	FiscalQuarter
)

type Quarter struct {
	calendar    calendar.Calendar
	quarterType QT
	year        int
	quarter     calendar.Q
	startDate   date.Date
}

func NewQuarter(cal calendar.Calendar, year int, qtr calendar.Q, qt QT) Quarter {
	switch qt {
	case CalendarQuarter:
		weekNo := ag7if.ComputeQuarterStartWeek(qtr)
		week := NewFromISOWeek(cal, year, weekNo)

		return Quarter{
			calendar:    cal,
			quarterType: qt,
			year:        year,
			quarter:     qtr,
			startDate:   week.Monday().Date(),
		}
	case FiscalQuarter:
		return Quarter{
			calendar:    cal,
			quarterType: qt,
			year:        year,
			quarter:     qtr,
			startDate:   cap.ComputeFiscalQuarterStartDate(year, qtr),
		}
	default:
		panic(errors.Errorf("invalid quarter type value: %d", qt))
	}
}

func (q Quarter) Next() calendar.Quarter {
	switch q.quarter {
	case calendar.Q1:
		return NewQuarter(q.calendar, q.year, calendar.Q2, q.quarterType)
	case calendar.Q2:
		return NewQuarter(q.calendar, q.year, calendar.Q3, q.quarterType)
	case calendar.Q3:
		return NewQuarter(q.calendar, q.year, calendar.Q4, q.quarterType)
	case calendar.Q4:
		return NewQuarter(q.calendar, q.year+1, calendar.Q1, q.quarterType)
	default:
		panic(errors.Errorf("invalid value for quarter: %d", q.quarter))
	}
}

func (q Quarter) StartDate() date.Date {
	return q.startDate
}

func (q Quarter) FirstWeek() calendar.Week {
	return NewWeek(q.calendar, q.startDate)
}

func (q Quarter) Short() string {
	return q.quarter.String()
}

func (q Quarter) String() string {
	switch q.quarterType {
	case CalendarQuarter:
		return fmt.Sprintf("CY%d%s", q.year, q.quarter)
	case FiscalQuarter:
		return fmt.Sprintf("FY%d%s", q.year, q.quarter)
	default:
		panic(errors.Errorf("invalid value for quarterType: %d", q.quarterType))
	}
}

func (q Quarter) Full() string {
	switch q.quarterType {
	case CalendarQuarter:
		return fmt.Sprintf("CY%d, %s Quarter", q.year, humanize.Ordinal(int(q.quarter)))
	case FiscalQuarter:
		return fmt.Sprintf("FY%d, %s Quarter", q.year, humanize.Ordinal(int(q.quarter)))
	default:
		panic(errors.Errorf("invalid value for quarterType: %d", q.quarterType))
	}
}
