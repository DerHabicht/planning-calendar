package cap

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
)

func ComputeQuarter(d date.Date) calendar.Q {
	switch d.Month() {
	case time.October, time.November, time.December:
		return calendar.Q1
	case time.January, time.February, time.March:
		return calendar.Q2
	case time.April, time.May, time.June:
		return calendar.Q3
	case time.July, time.August, time.September:
		return calendar.Q4
	default:
		panic(errors.Errorf("invalid month value: %d", d.Month()))
	}
}

func ComputeFiscalQuarterStartDate(fy int, q calendar.Q) date.Date {
	switch q {
	case calendar.Q1:
		return date.New(fy-1, time.October, 1)
	case calendar.Q2:
		return date.New(fy, time.January, 1)
	case calendar.Q3:
		return date.New(fy, time.April, 1)
	case calendar.Q4:
		return date.New(fy, time.July, 1)
	default:
		panic(errors.Errorf("invalid value for quarter: %d", q))
	}
}
