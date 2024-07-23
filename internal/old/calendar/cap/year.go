package cap

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

type FiscalYear struct {
	year      int
	startDate date.Date
	endDate   date.Date
}

func NewFiscalYear(year int) FiscalYear {
	sd := date.New(year-1, time.October, 1)
	ed := date.New(year, time.October, 1)

	startDate := calendar.ComputeNearestMonday(sd)
	endDate := calendar.ComputeNearestMonday(ed)

	return FiscalYear{
		year:      year,
		startDate: startDate,
		endDate:   endDate,
	}
}

func NewFiscalYearFromDate(d date.Date) FiscalYear {
	switch d.Month() {
	case time.November, time.December:
		return NewFiscalYear(d.Year() + 1)
	case time.October:
		fy := NewFiscalYear(d.Year() + 1)
		if !fy.IsWithinFiscalYear(d) {
			return NewFiscalYear(d.Year())
		}

		return fy
	case time.September:
		fy := NewFiscalYear(d.Year())
		if !fy.IsWithinFiscalYear(d) {
			return NewFiscalYear(d.Year() + 1)
		}

		return fy
	default:
		return NewFiscalYear(d.Year())
	}
}

func (fy FiscalYear) IsWithinFiscalYear(d date.Date) bool {
	return d.Equal(fy.startDate) || (d.After(fy.startDate) && d.Before(fy.endDate))
}

func (fy FiscalYear) FyWeek(d date.Date) (int, error) {
	if !fy.IsWithinFiscalYear(d) {
		return -1, fmt.Errorf("%s does not fall within FY%d", d.FormatISO(4), fy.year)
	}

	return (d.Sub(fy.startDate) / 7) + 1, nil
}
