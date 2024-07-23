package calendar

import (
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
)

type Calendar interface {
	FiscalYear() int
	JulianPeriod() int
	SolsticeTable() SolsticeTable
	FirstWeek() Week
}

type SolsticeTable interface {
	IsSolstice(date date.Date) Solstice
	FirstWinterSolstice() date.Date
	VernalEquinox() date.Date
	SummerSolstice() date.Date
	AutumnalEquinox() date.Date
	SecondWinterSolstice() date.Date
}

type Trimester interface {
	Year() int
	Trimester() T
	Next() Trimester
	StartDate() date.Date
	EndDate() date.Date
	FirstWeek() Week
	Short() string
	String() string
	Full() string
}

type Quarter interface {
	Year() int
	Quarter() Q
	Next() Quarter
	StartDate() date.Date
	EndDate() date.Date
	FirstWeek() Week
	Short() string
	String() string
	Full() string
}

type Month interface {
	Year() int
	Month() time.Month
	FirstWeek() Week
	Short() string
	String() string
}

type Sprint interface {
	Weeks() []Week
	Next() Sprint
	Short() string
	String() string
	Full() string
}

type Week interface {
	Trimester() Trimester
	FyQuarter() Quarter
	CyQuarter() Quarter
	Sprint() Sprint
	ISOWeek() (int, int, cards.Card)
	Next() Week
	Monday() Day
	Short() string
	String() string
	Full() string
}

type Day interface {
	Date() date.Date
	ISODate() string
	IsHoliday() (bool, bool, Holiday)
	IsSolstice() Solstice
	OrdinalDay() int
	MJD() int
	Sunrise() time.Time
	Sunset() time.Time
	Next() Day
}

type Holiday interface {
	Occurs(year int) (time.Time, time.Time)
	String() string
	FullName() string
}
