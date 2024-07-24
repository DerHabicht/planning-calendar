package plancal

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/cap"
)

type Trimester struct {
	calendar  calendar.Calendar
	trimester calendar.T
	year      int
	startDate date.Date
}

func NewTrimester(cal calendar.Calendar, year int, trimester calendar.T) Trimester {
	startDate := cap.ComputeTrimesterStartDate(year, trimester)

	return Trimester{
		calendar:  cal,
		trimester: trimester,
		year:      year,
		startDate: startDate,
	}
}

func (t Trimester) Next() calendar.Trimester {
	switch t.trimester {
	case calendar.T1:
		return NewTrimester(t.calendar, t.year, calendar.T2)
	case calendar.T2:
		return NewTrimester(t.calendar, t.year, calendar.T3)
	case calendar.T3:
		return NewTrimester(t.calendar, t.year+1, calendar.T1)
	default:
		panic(errors.Errorf("invalid trimester value: %d", t.trimester))
	}
}

func (t Trimester) StartDate() date.Date {
	return t.startDate
}

func (t Trimester) Short() string {
	return t.trimester.String()
}

func (t Trimester) String() string {
	return fmt.Sprintf("FY%d%s", t.year, t.trimester)
}

func (t Trimester) Full() string {
	return fmt.Sprintf("FY%d, %s Trimester", t.year, humanize.Ordinal(int(t.trimester)))
}
