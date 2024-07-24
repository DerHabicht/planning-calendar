package solstice_table

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"

	"github.com/derhabicht/planning-calendar/calendar"
)

type SolsticeTable struct {
	firstWinterSolstice  time.Time
	vernalEquinox        time.Time
	summerSolstice       time.Time
	autumnalEquinox      time.Time
	secondWinterSolstice time.Time
}

func NewSolsticeTable(fy int) *SolsticeTable {
	winter1 := julian.JDToTime(solstice.December(fy - 1))
	vernal := julian.JDToTime(solstice.March(fy))
	summmer := julian.JDToTime(solstice.June(fy))
	autumnal := julian.JDToTime(solstice.September(fy))
	winter2 := julian.JDToTime(solstice.December(fy))

	return &SolsticeTable{
		firstWinterSolstice:  winter1,
		vernalEquinox:        vernal,
		summerSolstice:       summmer,
		autumnalEquinox:      autumnal,
		secondWinterSolstice: winter2,
	}
}

func (s *SolsticeTable) IsSolstice(date date.Date) calendar.Solstice {
	switch date.Month() {
	case time.March:
		if date == calendar.TimeToLocalDate(s.vernalEquinox) {
			return calendar.VernalEquinox
		}
	case time.June:
		if date == calendar.TimeToLocalDate(s.summerSolstice) {
			return calendar.SummerSolstice
		}
	case time.September:
		if date == calendar.TimeToLocalDate(s.autumnalEquinox) {
			return calendar.AutumnalEquinox
		}
	case time.December:
		if date == calendar.TimeToLocalDate(s.firstWinterSolstice) || date == calendar.TimeToLocalDate(s.secondWinterSolstice) {
			return calendar.WinterSolstice
		}
	default:
		break
	}
	return calendar.NoSolstice
}

func (s *SolsticeTable) FirstWinterSolstice() date.Date {
	return calendar.TimeToLocalDate(s.firstWinterSolstice)
}

func (s *SolsticeTable) VernalEquinox() date.Date {
	return calendar.TimeToLocalDate(s.vernalEquinox)
}

func (s *SolsticeTable) SummerSolstice() date.Date {
	return calendar.TimeToLocalDate(s.summerSolstice)
}

func (s *SolsticeTable) AutumnalEquinox() date.Date {
	return calendar.TimeToLocalDate(s.autumnalEquinox)
}

func (s *SolsticeTable) SecondWinterSolstice() date.Date {
	return calendar.TimeToLocalDate(s.secondWinterSolstice)
}
