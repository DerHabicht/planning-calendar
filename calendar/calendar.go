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
	HolidayCalendar() HolidayCalendar
	FirstMonth() Month
}

type Trimester interface {
	Next() Trimester
	StartDate() date.Date
	Short() string
	String() string
	Full() string
}

type Quarter interface {
	Next() Quarter
	StartDate() date.Date
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
	Next() Month
}

type Sprint interface {
	FirstWeek() (Week, int)
	Short() string
	String() string
	Full() string
}

type Week interface {
	Trimester() Trimester
	FiscalQuarter() Quarter
	Quarter() Quarter
	Sprint() Sprint
	FyWeek() (int, int)
	ISOWeek() (int, int, cards.Card)
	Next() Week
	Monday() Day
	String() string
}

type Day interface {
	Date() date.Date
	ISODate() string
	IsHoliday() (bool, bool, Holiday)
	IsSolstice() Solstice
	MoonPhase() MoonPhase
	OrdinalDay() int
	MJD() int
	Sunrise() time.Time
	Sunset() time.Time
	Next() Day
}

type SolsticeTable interface {
	IsSolstice(date date.Date) Solstice
	FirstWinterSolstice() date.Date
	VernalEquinox() date.Date
	SummerSolstice() date.Date
	AutumnalEquinox() date.Date
	SecondWinterSolstice() date.Date
}

type HolidayCalendar interface {
	IsHoliday(date.Date) (bool, bool, Holiday)
	Holidays() []Holiday
}

type Holiday interface {
	Occurs(year int) (date.Date, date.Date)
	String() string
	FullName() string
}
