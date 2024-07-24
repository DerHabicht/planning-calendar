package cap

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
)

func ComputeTrimester(d date.Date) calendar.T {
	switch d.Month() {
	case time.October, time.November, time.December, time.January:
		return calendar.T1
	case time.February, time.March, time.April, time.May:
		return calendar.T2
	case time.June, time.July, time.August, time.September:
		return calendar.T3
	default:
		panic(errors.Errorf("invalid month value: %d", d.Month()))
	}
}

func ComputeTrimesterStartMonth(trimester calendar.T) time.Month {
	switch trimester {
	case calendar.T1:
		return time.October
	case calendar.T2:
		return time.February
	case calendar.T3:
		return time.June
	default:
		panic(errors.Errorf("invalid trimester value: %d", trimester))
	}
}

func ComputeTrimesterStartDate(year int, trimester calendar.T) date.Date {
	month := ComputeTrimesterStartMonth(trimester)

	if trimester == calendar.T1 {
		year -= 1
	}

	return date.New(year, month, 1)
}
