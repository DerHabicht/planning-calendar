package calendar

import (
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

type Calendar interface {
	FiscalYear() int
	SolsticeTable() SolsticeTable
	Params() map[string]string
	CurrentMonth() Month
	NextMonth() Month
	CurrentWeek() Week
	NextWeek() Week
}

type Quarter interface{}

type Month interface {
	Year() int
	Month() time.Month
	String() string
}

type Week interface {
	Next() Week
	Reset() Week
	CurrentDay() Day
	NextDay() Day
	Params() map[string]string
}

type Day interface {
	Date() date.Date
	String() string
	IsHoliday() (bool, string)
	Sunrise() time.Time
	Sunset() time.Time
}

type Sprint interface{}

type Holiday interface{}

type SolsticeTable interface {
	FiscalYear() int
	FirstWinterSolstice() time.Time
	VernalEquinox() time.Time
	SummerSolstice() time.Time
	AutumnalEquinox() time.Time
	SecondWinterSolstice() time.Time
	IsSolstice(d date.Date) (bool, calendar.Solstice)
}
