package cap

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
)

type T int

const (
	T1 T = iota + 1
	T2
	T3
)

func (f T) String() string {
	switch f {
	case T1:
		return "T1"
	case T2:
		return "T2"
	case T3:
		return "T3"
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

type Trimester struct {
	Trimester  T
	FiscalYear int
}

func NewTrimester(d date.Date) Trimester {
	switch d.Month() {
	case time.October, time.November, time.December, time.January:
		if d.Month() == time.January {
			return Trimester{T1, d.Year()}
		}
		return Trimester{T1, d.Year() + 1}
	case time.February, time.March, time.April, time.May:
		return Trimester{T2, d.Year()}
	case time.June, time.July, time.August, time.September:
		return Trimester{T3, d.Year()}
	default:
		panic(fmt.Errorf("%d is not a valid month", d.Month()))
	}
}

func (f Trimester) StartDate() date.Date {
	switch f.Trimester {
	case T1:
		return date.New(f.FiscalYear-1, time.October, 1)
	case T2:
		return date.New(f.FiscalYear, time.February, 1)
	case T3:
		return date.New(f.FiscalYear, time.June, 1)
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f Trimester) NextTrimester() Trimester {
	switch f.Trimester {
	case T1:
		return Trimester{
			Trimester:  T2,
			FiscalYear: f.FiscalYear,
		}
	case T2:
		return Trimester{
			Trimester:  T3,
			FiscalYear: f.FiscalYear,
		}
	case T3:
		return Trimester{
			Trimester:  T1,
			FiscalYear: f.FiscalYear + 1,
		}
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f Trimester) String() string {
	return fmt.Sprintf("FY%d%s", f.FiscalYear, f.Trimester)
}
