package cap

import (
	"time"

	"github.com/fxtlabs/date"
)

func ComputeFiscalYear(d date.Date) int {
	if d.Month() == time.October || d.Month() == time.November || d.Month() == time.December {
		return d.Year() - 1
	}

	return d.Year()
}
